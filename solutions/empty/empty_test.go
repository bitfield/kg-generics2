package empty_test

import (
	"testing"

	"github.com/bitfield/empty"
)

func TestEmptyIsTrueForEmptySequence(t *testing.T) {
	t.Parallel()
	s := empty.Sequence[int]{}
	if !s.Empty() {
		t.Fatal("Empty(%v): want false, got true")
	}
}

func TestEmptyIsFalseForNonEmptySequence(t *testing.T) {
	t.Parallel()
	s := empty.Sequence[string]{"a", "b", "c"}
	if s.Empty() {
		t.Fatal("Empty(%v): want true, got false")
	}
}
