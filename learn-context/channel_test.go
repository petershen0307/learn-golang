package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_channel_always_allow(t *testing.T) {
	ch := make(chan int)
	close(ch)
	<-ch
	assert.True(t, true)
}

func Test_channel_always_block(t *testing.T) {
	var ch chan int
	select {
	case ch <- 1:
		assert.True(t, false)
	case <-ch:
		assert.True(t, false)
	default:
	}
	assert.True(t, true)
}

func Test_background_context_always_block(t *testing.T) {
	bkCtx := context.Background()
	select {
	case <-bkCtx.Done():
		assert.True(t, false)
	default:
	}
	assert.True(t, true)
}
