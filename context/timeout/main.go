package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	start := time.Now()

	if err := parentFunc(ctx); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("time passed %v\n", time.Since(start))
	/*
		error: parent error: childFunc error: heavyFunc error: context deadline exceeded
		time passed 1.001596379s
	*/
}

func parentFunc(ctx context.Context) error {
	d := time.Second
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	if err := childFunc(ctx); err != nil {
		return fmt.Errorf("parent error: %w", err)
	}

	return nil
}

func childFunc(ctx context.Context) error {
	// Deadline of parent is 2 sec.
	d := 2 * time.Second
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	if err := heavyFunc(ctx); err != nil {
		return fmt.Errorf("childFunc error: %w", err)
	}

	return nil
}

func heavyFunc(ctx context.Context) error {
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("finish calculation")
		return nil
	case <-ctx.Done():
		return fmt.Errorf("heavyFunc error: %w", ctx.Err())
	}
}

//
//func main() {
//	start := time.Now()
//
//	parent := context.Background()
//	ctx, cancel := context.WithTimeout(parent, time.Second)
//	defer cancel()
//
//	if err := parentFunc(ctx); err != nil {
//		fmt.Printf("error: %v\n", err)
//	}
//
//	fmt.Printf("time passed %v\n", time.Since(start))
//}
