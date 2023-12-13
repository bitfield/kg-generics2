package channel_test

import (
	"sync"
	"testing"

	"github.com/bitfield/channel"
)

func TestSendFollowedByReceiveGivesOriginalValue(t *testing.T) {
	t.Parallel()
	c := channel.New[int](1)
	want := 99
	c.Send(want)
	got := c.Receive()
	if want != got {
		t.Fatalf("want %d received, got %d", want, got)
	}
}

func TestSendsReturns3AfterThreeSendOperations(t *testing.T) {
	t.Parallel()
	c := channel.New[float64](3)
	c.Send(1.0)
	c.Send(2.0)
	c.Send(3.0)
	want := uint64(3)
	got := c.Sends()
	if want != got {
		t.Fatalf("want %d sends, got %d", want, got)
	}
}

func TestReceivesReturns2AfterTwoSendOperations(t *testing.T) {
	t.Parallel()
	c := channel.New[struct{}](1)
	c.Send(struct{}{})
	_ = c.Receive()
	c.Send(struct{}{})
	_ = c.Receive()
	want := uint64(2)
	got := c.Receives()
	if want != got {
		t.Fatalf("want %d receives, got %d", want, got)
	}
}

func TestChannelHandlesConcurrentSendsAndReceives(t *testing.T) {
	t.Parallel()
	c := channel.New[string](10)
	want := uint64(100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := uint64(0); i < want; i++ {
			c.Send("hello")
			_ = c.Receives()
		}
		wg.Done()
	}()
	for i := uint64(0); i < want; i++ {
		_ = c.Receive()
		_ = c.Sends()
	}
	wg.Wait()
	got := c.Sends()
	if want != got {
		t.Errorf("want %d sends, got %d", want, got)
	}
	got = c.Receives()
	if want != got {
		t.Errorf("want %d receives, got %d", want, got)
	}
}
