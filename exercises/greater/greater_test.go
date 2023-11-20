package greater_test

import (
	"testing"

	"github.com/bitfield/greater"
)

type MyInt int

func (m MyInt) Greater(v MyInt) bool {
	return m > v
}

func TestIsGreater_IsTrueFor2And1(t *testing.T) {
	t.Parallel()
	if !greater.IsGreater(MyInt(2), MyInt(1)) {
		t.Fatalf("IsGreater(2, 1): want true, got false")
	}
}

func TestIsGreater_IsFalseFor1And2(t *testing.T) {
	t.Parallel()
	if greater.IsGreater(MyInt(1), MyInt(2)) {
		t.Fatalf("IsGreater(1, 2): want false, got true")
	}
}
