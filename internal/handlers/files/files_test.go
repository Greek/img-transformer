package files

import (
	"slices"
	"testing"
)

func TestParseTransformation(t *testing.T) {
	path := "/testbucket1/testfile1.png=w32,h32,r16"
	expected := []string{"w32", "h32", "r16"}

	output := parseTransformation(path)
	if !slices.Equal(output, expected) {
		t.Errorf(`parseTransformation("%q") = %q, want %q`, path, output, expected)
	}
}
