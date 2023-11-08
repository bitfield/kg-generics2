package funcmap_test

import (
	"testing"
	"unicode"

	"github.com/bitfield/funcmap"

	"github.com/google/go-cmp/cmp"
)

func TestFuncMap_AppliesDoubleTo2AndReturns4(t *testing.T) {
	t.Parallel()
	fm := funcmap.FuncMap[int, int]{
		"double": func(i int) int {
			return i * 2
		},
		"addOne": func(i int) int {
			return i + 1
		},
	}
	want := 4
	got := fm.Apply("double", 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFuncMap_AppliesUpperToUppercaseInputAndReturnsTrue(t *testing.T) {
	t.Parallel()
	fm := funcmap.FuncMap[rune, bool]{
		"upper": unicode.IsUpper,
		"lower": unicode.IsLower,
	}
	got := fm.Apply("upper", 'A')
	if !got {
		t.Fatal("upper('A'): want true, got false")
	}
}
