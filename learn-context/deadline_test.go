package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_timeout(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	now := time.Now()
	<-ctx.Done()
	diff := time.Since(now)
	log.Println("diff=", diff.Seconds())
	assert.GreaterOrEqual(t, diff, time.Second)
}

func Test_cancel_before_timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	now := time.Now()
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	<-ctx.Done()
	diff := time.Since(now)
	log.Println("diff=", diff.Seconds())
	assert.GreaterOrEqual(t, diff, time.Second)
}

func Test_deadline(t *testing.T) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	now := time.Now()
	<-ctx.Done()
	diff := time.Since(now)
	log.Println("diff=", diff.Seconds())
	assert.GreaterOrEqual(t, diff, time.Second)
}

func Test_cancel_before_deadline(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	now := time.Now()
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	<-ctx.Done()
	diff := time.Since(now)
	log.Println("diff=", diff.Seconds())
	assert.GreaterOrEqual(t, diff, time.Second)
}
