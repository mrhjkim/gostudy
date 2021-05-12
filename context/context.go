package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func longFunc() string {
	<-time.After(time.Second * 3)
	return "Success"
}

func longFuncWithCtx(ctx context.Context) (string, error) {
	done := make(chan string)

	go func() {
		done <- longFunc()
	}()

	select {
	case result := <-done:
		return result, nil
	case <-ctx.Done():
		return "Fail", ctx.Err()

	}
}
func runWithoutCancel() {
	ctx, _ := context.WithCancel(context.Background())

	result, err := longFuncWithCtx(ctx)
	if err != nil {
		log.Printf("FAIL : Context cancelled with err: %v", err)
	} else {
		log.Printf("Context not cancelled with result: %v", result)
	}
}
func runWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		cancel()
	}()

	result, err := longFuncWithCtx(ctx)
	if err != nil {
		log.Printf("Context cancelled with err: %v", err)
	} else {
		log.Printf("FAIL : Context not cancelled with result: %v", result)
	}

}

func runWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := longFuncWithCtx(ctx)
	if err != nil {
		log.Printf("Context timeout with err: %v", err)
	} else {
		log.Printf("FAIL : Context with result: %v", result)
	}
}
func myFunc(ctx context.Context) error {
	if v := ctx.Value("current_user"); v != nil {
		u, ok := v.(string)
		if !ok {
			return errors.New("Type error")
		}
		log.Printf("user:%s", u)
		return nil
	}
	return errors.New("No user")
}
func ctxWithValue() {
	ctx := context.WithValue(context.Background(), "current_user", "A")
	myFunc(ctx)
}
func main() {
	runWithCancel()
	runWithoutCancel()
	runWithTimeout()
	ctxWithValue()
}
