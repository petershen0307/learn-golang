package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func query(ctx context.Context, id int, result chan int) {
	timer := time.NewTimer(2 * time.Second)
	select {
	case <-ctx.Done():
		log.Println("id=", id, "query cancel!")
	case <-timer.C:
		log.Println("id=", id, "query finish!")
		result <- id
	}
}

func queryTask(ctx context.Context, queryCount int) chan int {
	resultChan := make(chan int, queryCount)
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
	log.Println("result length", len(resultChan))
	for i := 0; i != queryCount; i++ {
		log.Println("result", <-resultChan)
	}
}

func Test_cancel_query(t *testing.T) {
	const queryCount = 10
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	resultChan := queryTask(cancelCtx, queryCount)

	time.Sleep(1 * time.Second)
	cancelFunc()
	time.Sleep(1 * time.Second)
	assert.Equal(t, 0, len(resultChan))
	log.Println("result length", len(resultChan))
}
