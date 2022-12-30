package context

// func Background() Context
// Returns an empty Context that’s never cancelled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

// func TODO() Context
// Also provides an empty Context, but it’s intended to be used as a placeholder when it’s unclear which Context to use or when a parent Context is not yet avail‐ able.

// func WithDeadline(Context, time.Time) (Context, CancelFunc)
// Accepts a specific time at which the Context will be cancelled and the Done chan‐ nel will be closed.

// func WithTimeout(Context, time.Duration) (Context, CancelFunc)
// Accepts a duration after which the Context will be cancelled and the Done chan‐ nel will be closed.

// func WithCancel(Context) (Context, CancelFunc)
// Unlike the previous functions, WithCancel accepts nothing, and only returns a function that can be called to explicitly cancel the Context.

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestBackGroundContext(t *testing.T) {
	ctx := context.Background()
	res := backGroundContext(ctx)
	t.Log(res)

}

func TestWithValueContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", 1)
	res := withValueContext(ctx)
	if res != "1" {
		t.Error(`res should be "1"`)
	}
	t.Log(res)
}


func TestWithCancelContext(t *testing.T) {

}

func backGroundContext(ctx context.Context) string {
	return "this msg is from sample service"
}

func withValueContext(ctx context.Context) string {

	return fmt.Sprint(ctx.Value("key"))
}

func withCancelContext(ctx context.Context, delay time.Duration) string {
	withCancelContextChild01 := func(ctx context.Context, delay time.Duration) string {
		time.Sleep(delay * time.Second)
		return "return from withCancelContextChild01"
	}

	withCancelContextChild02 := func(ctx context.Context, delay time.Duration) string {
		time.Sleep(delay * time.Second)
		return "return from withCancelContextChild02"
	}

	withCancelContextChild01(ctx, delay)
	withCancelContextChild02(ctx, delay)

	time.Sleep(delay * time.Second)
	return "return from withCancelContext"
}
