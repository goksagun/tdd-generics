package main

func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	if Contains([]int{1, 2, 3}, 10) {
		println("OK")
	} else {
		println("KO")
	}
}
