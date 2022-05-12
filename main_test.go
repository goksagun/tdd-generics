package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("test it contains integer item", func(t *testing.T) {
		got := Contains([]int{1, 2, 3}, 1)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test it not contains integer item", func(t *testing.T) {
		got := Contains([]int{1, 2, 3}, 9)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestContainsString(t *testing.T) {
	t.Run("test it contains string item", func(t *testing.T) {
		got := ContainsString([]string{"foo", "bar", "baz"}, "foo")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test it not contains string item", func(t *testing.T) {
		got := ContainsString([]string{"foo", "bar", "baz"}, "foobar")
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
