# Using Generics Array Contains Function

A generic approach to check if an element is present in an array/slice

### Requirements
- Golang 1.18+

## Approach
Our intention should be clear. So what is our goal? It’s pretty simple. We want a function that is passed an array and a value and returns true if the array contains the value and false if the value is missing.

## Solution
One solution would be to create a function where we set the property types statically. In the function, we simply iterate over the array and see if any of the values in the array are equal to our element.

*first test...*

```go
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
```

*...implement simple working code*

```go
package main

func Contains(items []int, item int) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func main()  {
	if Contains([]int{1, 2, 3}, 10) {
		println("OK")	
	} else {
		println("KO")
	}	
}
```

*run tests*

```sh
go test
```

```sh
❯ go test
PASS
ok      example 0.196s
```

It works fine.

*...one more thing add fail test*

```go
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
```

*run tests again*

```sh
go test
```

```sh
❯ go test
PASS
ok      example 0.196s
```

## Contains string

Imagine that for many arrays with different types we need to check if an array contains a string. We need to create a new function for string type.

*first test...*

```go
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
```

*...implement simple working code*

```go
func ContainsString(items []string, item string) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
```

*run tests*

```sh
go test
```

```sh
❯ go test
PASS
ok      example 0.103s
```

## Refactoring... extract assertion method

It's time to refactor and remove duplicates. Move assertion to new `assertCorrectMessage` method.

```go
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
```

*run tests*

```sh
go test
```

```sh
❯ go test
PASS
ok      example 0.109s
```

Allt he test passes and everytihng seems ok.

## What about contains float?

*first test...*

```go
func TestContainsFloat(t *testing.T) {
	t.Run("test it contains float item", func(t *testing.T) {
		got := ContainsFloat([]float32{1.15, 2.75, 3.99}, 1.15)
		want := true

		assertCorrectMessage(got, want, t)
	})
	t.Run("test it not contains float item", func(t *testing.T) {
		got := ContainsFloat([]float32{1.15, 2.75, 3.99}, 9.85)
		want := false

		assertCorrectMessage(got, want, t)
	})
}
```

*...implement simple working code*

```go
func ContainsFloat(items []float32, item float32) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
```

*run tests*

```sh
go test
```

```sh
❯ go test
PASS
ok      example 0.106s
```

It works fine. We have only one problem with this approach. Imagine that for many arrays with different types we need to check if an array contains a value. To archive this, we need to create a new function for each type. This leads to a very verbose and duplicate code base. 

Fortunately for us, Golang introduced a number of different generic types a few months ago. In the next chapter, we can use one of them to make our solution a little bit more flexible.

## Generics magic..!

For our use case, we need the “comparable” generic type. (If you want to know more about generics, I can recommend  the following [article][1].)

As you can see below, we need to write the name of our generic T in brackets right after the function name and assign it the type “comparable”. We also give both arguments the type T. This way, when we call the function, we can specify which type of array and element we want to compare.

*first test...*

```go
func TestContains(t *testing.T) {
	t.Run("test it contains integer item", func(t *testing.T) {
		got := Contains[int]([]int{1, 2, 3}, 1)
		want := true

		assertCorrectMessage(got, want, t)
	})
	t.Run("test it not contains integer item", func(t *testing.T) {
		got := Contains[int]([]int{1, 2, 3}, 9)
		want := false

		assertCorrectMessage(got, want, t)
	})
	t.Run("test it contains string item", func(t *testing.T) {
		got := Contains[string]([]string{"foo", "bar", "baz"}, "foo")
		want := true

		assertCorrectMessage(got, want, t)
	})
	t.Run("test it not contains string item", func(t *testing.T) {
		got := Contains[string]([]string{"foo", "bar", "baz"}, "foobar")
		want := false

		assertCorrectMessage(got, want, t)
	})
	t.Run("test it contains float item", func(t *testing.T) {
		got := Contains[float32]([]float32{1.15, 2.75, 3.99}, 1.15)
		want := true

		assertCorrectMessage(got, want, t)
	})
	t.Run("test it not contains float item", func(t *testing.T) {
		got := Contains[float32]([]float32{1.15, 2.75, 3.99}, 9.85)
		want := false

		assertCorrectMessage(got, want, t)
	})
}
```

*...implement simple working code*

```go
func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
```

To use it, we have to make the following function call: `Contains[string](["bar", "foo"], "foo")`. The word “string” inside the brackets before the parenthesis is where we specify the type of our array and value.

*run tests*

```sh
go test
```

```sh
❯ go test
PASS
ok      example 0.106s
```

[1]: https://bitfieldconsulting.com/golang/generics
