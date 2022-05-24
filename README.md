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
