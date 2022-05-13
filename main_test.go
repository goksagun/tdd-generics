package main

import (
	"testing"
)

func assertCorrectMessage(got bool, want bool, t *testing.T) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestContains(t *testing.T) {
	t.Run("test it contains integer item", func(t *testing.T) {
		got := Contains([]int{1, 2, 3}, 1)
		want := true

		assertCorrectMessage(got, want, t)
	})
	t.Run("test it not contains integer item", func(t *testing.T) {
		got := Contains([]int{1, 2, 3}, 9)
		want := false

		assertCorrectMessage(got, want, t)
	})
}

func TestContainsString(t *testing.T) {
	t.Run("test it contains string item", func(t *testing.T) {
		got := ContainsString([]string{"foo", "bar", "baz"}, "foo")
		want := true

		assertCorrectMessage(got, want, t)
	})
	t.Run("test it not contains string item", func(t *testing.T) {
		got := ContainsString([]string{"foo", "bar", "baz"}, "foobar")
		want := false

		assertCorrectMessage(got, want, t)
	})
}
