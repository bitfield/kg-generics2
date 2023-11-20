package compose_test

import (
	"strings"
	"testing"

	"github.com/bitfield/compose"
)

func isOdd(p int) bool {
	return p%2 != 0
}

func next(p int) int {
	return p + 1
}

func TestComposeAppliesFuncsToIntInReverseOrder(t *testing.T) {
	t.Parallel()
	if compose.Compose(isOdd, next, 1) {
		t.Fatal("want false for isOdd(next(1)), got true")
	}
}

func TestComposeAppliesFuncsToStringInReverseOrder(t *testing.T) {
	t.Parallel()
	input := "HeLlO, wOrLd"
	want := "hello, world"
	got := compose.Compose(strings.ToLower, strings.ToUpper, input)
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
