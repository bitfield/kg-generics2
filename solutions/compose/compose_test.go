package compose_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/bitfield/compose"
)

func double(p int) int {
	return p * 2
}

func addOne(p int) int {
	return p + 1
}

func TestComposeAppliesFuncsInReverseOrder(t *testing.T) {
	t.Parallel()
	want := 4
	got := compose.Compose(double, addOne, 1)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestComposeHandlesDifferentFuncSignatures(t *testing.T) {
	t.Parallel()
	input := "HeLlO, wOrLd"
	want := 12
	got := compose.Compose(utf8.RuneCountInString, strings.ToUpper, input)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func first[E any](s []E) E {
	return s[0]
}

func last[E any](s []E) E {
	return s[len(s)-1]
}

func TestComposeAppliesFuncsToSliceInReverseOrder(t *testing.T) {
	t.Parallel()
	input := [][]int{{1, 2, 3}}
	want := 1
	got := compose.Compose(first[int], last[[]int], input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestComposeAppliesFuncsToSliceInReverseOrder2(t *testing.T) {
	t.Parallel()
	input := [][]int{{1, 2, 3}}
	want := 3
	got := compose.Compose(last[int], first[[]int], input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
