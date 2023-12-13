package empty_test

import (
	"testing"

	"github.com/bitfield/empty"
)

func TestEmptyIsTrueForEmptySequence(t *testing.T) {
	t.Parallel()
	s := empty.Sequence[int]{}
	if !s.Empty() {
		t.Fatalf("Empty(%v): want true, got false", s)
	}
}

func TestEmptyIsFalseForNonEmptySequence(t *testing.T) {
	t.Parallel()
	s := empty.Sequence[string]{"a", "b", "c"}
	if s.Empty() {
		t.Fatalf("Empty(%v): want false, got true", s)
	}
}
