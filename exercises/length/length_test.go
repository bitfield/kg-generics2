package length_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/bitfield/length"
)

func TestLenOfGroupIs2WhenItContains2Elements(t *testing.T) {
	t.Parallel()
	want := 2
	got := length.Len([]int{1, 2})
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
