package channel

import "sync/atomic"

type Channel[T any] struct {
	ch              chan T
	sends, receives atomic.Uint64
}

func New[T any](length int) *Channel[T] {
	return &Channel[T]{
		ch: make(chan T, length),
	}
}

func (c *Channel[T]) Send(v T) {
	c.ch <- v
	c.sends.Add(1)
}

func (c *Channel[T]) Receive() T {
	v := <-c.ch
	c.receives.Add(1)
	return v
}

func (c *Channel[T]) Sends() uint64 {
	return c.sends.Load()
}

func (c *Channel[T]) Receives() uint64 {
	return c.receives.Load()
}
