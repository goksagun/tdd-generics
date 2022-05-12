package main

func Contains(items []int, item int) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func ContainsString(items []string, item string) bool {
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