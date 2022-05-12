package main

import (
	"testing"
)

func TestContains(t *testing.T) {	
	got := Contains([]int{1,2,3}, 1)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
