package stringy_test

import (
	"bytes"
	"testing"

	"github.com/bitfield/stringy"
)

type greeting struct{}

func (greeting) String() string {
	return "Howdy!"
}

func TestStringifyTo_PrintsToSuppliedWriter(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	stringy.StringifyTo[greeting](buf, greeting{})
	want := "Howdy!\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
