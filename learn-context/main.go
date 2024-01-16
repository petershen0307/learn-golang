package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	run1(ctx)
	time.Sleep(50 * time.Millisecond)
	cancel()
	time.Sleep(50 * time.Millisecond)
}

func run1(ctx context.Context) {
	go func() {
		<-ctx.Done()
		defer fmt.Println("run1!")
	}()
	run2(ctx)
}

func run2(ctx context.Context) {
	go func() {
		<-ctx.Done()
		defer fmt.Println("run2!")
	}()
	run3(ctx)
}

func run3(ctx context.Context) {
	go func() {
		<-ctx.Done()
		defer fmt.Println("run3!")
	}()
}
