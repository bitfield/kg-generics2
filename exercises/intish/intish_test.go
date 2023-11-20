package intish_test

import (
	"testing"

	"github.com/bitfield/intish"
)

type MyInt int

func TestIsPositive_IsTrueFor1(t *testing.T) {
	t.Parallel()
	input := MyInt(1)
	if !intish.IsPositive(input) {
		t.Errorf("IsPositive(1): want true, got false")
	}
}

func TestIsPositive_IsFalseForNegative1(t *testing.T) {
	t.Parallel()
	input := MyInt(-1)
	if intish.IsPositive(input) {
		t.Errorf("IsPositive(-1): want false, got true")
	}
}

func TestIsPositive_IsFalseForZero(t *testing.T) {
	t.Parallel()
	input := MyInt(0)
	if intish.IsPositive(input) {
		t.Errorf("IsPositive(0): want false, got true")
	}
}
