package intish_test

import (
	"testing"

	"github.com/bitfield/intish"

	"github.com/google/go-cmp/cmp"
)

type MyInt int

func TestIsPositive_IsTrueFor1(t *testing.T) {
	t.Parallel()
	input := MyInt(1)
	want := true
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositive_IsFalseForNegative1(t *testing.T) {
	t.Parallel()
	input := MyInt(-1)
	want := false
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositive_IsFalseForZero(t *testing.T) {
	t.Parallel()
	input := MyInt(0)
	want := false
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
