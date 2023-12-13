package merge_test

import (
	"maps"
	"testing"

	"github.com/bitfield/merge"
)

func TestMergeCorrectlyMergesTwoMapsOfIntToBool(t *testing.T) {
	t.Parallel()
	inputs := []map[int]bool{
		{
			1: false,
			2: false,
			3: false,
		},
		{
			3: true,
			5: true,
		},
	}
	want := map[int]bool{
		1: false,
		2: false,
		3: true,
		5: true,
	}
	got := merge.Merge(inputs...)
	if !maps.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMergeCorrectlyMergesThreeMapsOfStringToAny(t *testing.T) {
	t.Parallel()
	inputs := []map[string]any{
		{
			"a": nil,
		},
		{
			"b": "hello, world",
			"c": 0,
		},
		{
			"a": 6 + 2i,
		},
	}
	want := map[string]any{
		"a": 6 + 2i,
		"b": "hello, world",
		"c": 0,
	}
	got := merge.Merge(inputs...)
	if !maps.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMergeHandlesDerivedTypes(t *testing.T) {
	t.Parallel()
	type menu map[int]string
	m1 := menu{1: "eggs"}
	m2 := menu{2: "beans"}
	want := menu{
		1: "eggs",
		2: "beans",
	}
	got := merge.Merge(m1, m2)
	if !maps.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
