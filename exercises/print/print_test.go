package print_test

import (
	"bytes"
	"testing"

	"github.com/bitfield/print"
)

func TestPrintAnythingTo_PrintsToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	print.PrintAnythingTo(buf, "Hello, world")
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
