package main

import (
	"testing"
)

func TestContains(t *testing.T) {	
	t.Run("test it contains integer item", func(t *testing.T) {
		got := Contains([]int{1,2,3}, 1)
		want := true
	
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
