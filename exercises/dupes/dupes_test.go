package dupes_test

import (
	"testing"

	"github.com/bitfield/dupes"
)

func TestDupesIsTrueGivenNonConsecutiveDuplicates(t *testing.T) {
	t.Parallel()
	s := []int{1, 2, 3, 1, 5}
	if !dupes.Dupes(s) {
		t.Errorf("Dupes(%v): want true, got false", s)
	}
}

func TestDupesIsFalseGivenNoDuplicates(t *testing.T) {
	t.Parallel()
	s := []string{"a", "b", "c"}
	if dupes.Dupes(s) {
		t.Errorf("Dupes(%v): want false, got true", s)
	}
}
