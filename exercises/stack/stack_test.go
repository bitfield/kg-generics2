package stack_test

import (
	"testing"

	"github.com/bitfield/stack"
)

func TestPushToEmptyStackGivesStackWithLength1(t *testing.T) {
	t.Parallel()
	s := stack.Stack[int]{}
	s.Push(0)
	if s.Len() != 1 {
		t.Fatalf("want stack length 1, got %d", s.Len())
	}
}

func TestPushPushPopPopLeavesStackEmpty(t *testing.T) {
	t.Parallel()
	s := stack.Stack[string]{}
	s.Push("a", "b")
	if s.Len() != 2 {
		t.Fatal("Push didn't add all values to stack")
	}
	s.Pop()
	want := "a"
	got, ok := s.Pop()
	if !ok {
		t.Fatal("Pop returned not ok on non-empty stack")
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestPopReturnsNotOKOnEmptyStack(t *testing.T) {
	t.Parallel()
	s := stack.Stack[int]{}
	_, ok := s.Pop()
	if ok {
		t.Fatal("Pop returned ok on empty stack")
	}
}
