package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetVariables(t *testing.T) {
	*prefix = "TEST_"
	os.Setenv("TEST_THIS_STUFF", "yes")
	os.Setenv("TEST_OTHER_STUFF", "sure")
	got := getVariables()
	want := map[string]string{
		"ThisStuff":  "yes",
		"OtherStuff": "sure",
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Incorrect variables exposed: %s", diff)
	}
}
