package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func query(ctx context.Context, id int, result chan string) {
	timer := time.NewTimer(2 * time.Second)
	select {
	case <-ctx.Done():
		s := fmt.Sprint("id=", id, "query cancel!")
		log.Println(s)
		result <- s
	case <-timer.C:
		s := fmt.Sprint("id=", id, "query finish!")
		log.Println(s)
		result <- s
	}
}

func queryTask(ctx context.Context, queryCount int) chan string {
	resultChan := make(chan string, queryCount)
	for i := 0; i != queryCount; i++ {
		go query(ctx, i, resultChan)
	}
	return resultChan
}

func Test_query(t *testing.T) {
	const queryCount = 10
	cancelCtx, _ := context.WithCancel(context.Background())

	resultChan := queryTask(cancelCtx, queryCount)

	time.Sleep(3 * time.Second)
	assert.Equal(t, queryCount, len(resultChan))
	for i := 0; i != queryCount; i++ {
		s := <-resultChan
		assert.Contains(t, s, "query finish!")
		log.Println("result", s)
	}
}

func Test_cancel_query(t *testing.T) {
	const queryCount = 10
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	resultChan := queryTask(cancelCtx, queryCount)

	time.Sleep(1 * time.Second)
	cancelFunc()
	time.Sleep(1 * time.Second)
	assert.Equal(t, queryCount, len(resultChan))
	for i := 0; i != queryCount; i++ {
		s := <-resultChan
		assert.Contains(t, s, "query cancel!")
		log.Println("result", s)
	}
}

func Test_spawn_context_one_layer_cancel_from_leaf(t *testing.T) {
	ctx1, _ := context.WithCancel(context.Background())
	ctx11, cancelFunc11 := context.WithCancel(ctx1)
	ctx12, cancelFunc12 := context.WithCancel(ctx1)
	resultChan := make(chan string, 2)

	go query(ctx11, 1, resultChan)
	go query(ctx12, 2, resultChan)
	cancelFunc11()
	cancelFunc12()

	time.Sleep(1 * time.Second)
	assert.Equal(t, 2, len(resultChan))
	for i := 0; i != 2; i++ {
		s := <-resultChan
		assert.Contains(t, s, "query cancel!")
		log.Println("result", s)
	}
}

func Test_spawn_context_one_layer_cancel_from_leaf_partial(t *testing.T) {
	ctx1, _ := context.WithCancel(context.Background())
	ctx11, cancelFunc11 := context.WithCancel(ctx1)
	ctx12, _ := context.WithCancel(ctx1)
	resultChan := make(chan string, 2)

	go query(ctx11, 1, resultChan)
	go query(ctx12, 2, resultChan)
	cancelFunc11()

	time.Sleep(1 * time.Second)
	s := <-resultChan
	assert.Contains(t, s, "query cancel!")
	log.Println("result", s)
	s = <-resultChan
	assert.Contains(t, s, "query finish!")
	log.Println("result", s)
}

func Test_spawn_context_one_layer_cancel_from_root(t *testing.T) {
	ctx1, cancelFunc1 := context.WithCancel(context.Background())
	ctx11, _ := context.WithCancel(ctx1)
	ctx12, _ := context.WithCancel(ctx1)
	resultChan := make(chan string, 2)

	go query(ctx11, 1, resultChan)
	go query(ctx12, 2, resultChan)
	cancelFunc1()

	time.Sleep(1 * time.Second)
	assert.Equal(t, 2, len(resultChan))
	for i := 0; i != 2; i++ {
		s := <-resultChan
		assert.Contains(t, s, "query cancel!")
		log.Println("result", s)
	}
}

func Test_spawn_context_multiple_layer_cancel_from_leaf(t *testing.T) {
	ctx1, _ := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx1)
	ctx3, cancelFunc3 := context.WithCancel(ctx2)
	resultChan := make(chan string, 2)

	go query(ctx2, 1, resultChan)
	go query(ctx3, 2, resultChan)
	cancelFunc3()

	time.Sleep(1 * time.Second)
	s := <-resultChan
	assert.Contains(t, s, "query cancel!")
	log.Println("result", s)
	s = <-resultChan
	assert.Contains(t, s, "query finish!")
	log.Println("result", s)
}

func Test_spawn_context_multiple_layer_cancel_from_middle(t *testing.T) {
	ctx1, _ := context.WithCancel(context.Background())
	ctx2, cancelFunc2 := context.WithCancel(ctx1)
	ctx3, _ := context.WithCancel(ctx2)
	resultChan := make(chan string, 2)

	go query(ctx2, 1, resultChan)
	go query(ctx3, 2, resultChan)
	cancelFunc2()

	time.Sleep(1 * time.Second)
	assert.Equal(t, 2, len(resultChan))
	for i := 0; i != 2; i++ {
		s := <-resultChan
		assert.Contains(t, s, "query cancel!")
		log.Println("result", s)
	}
}

func Test_spawn_context_multiple_layer_cancel_from_root(t *testing.T) {
	ctx1, cancelFunc1 := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx1)
	ctx3, _ := context.WithCancel(ctx2)
	resultChan := make(chan string, 2)

	go query(ctx2, 1, resultChan)
	go query(ctx3, 2, resultChan)
	cancelFunc1()

	time.Sleep(1 * time.Second)
	assert.Equal(t, 2, len(resultChan))
	for i := 0; i != 2; i++ {
		s := <-resultChan
		assert.Contains(t, s, "query cancel!")
		log.Println("result", s)
	}
}
