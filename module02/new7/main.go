package main

import "fmt"

func contains(x string, a ...string) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func getmax(a ...int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	A := []string{"a", "b"}
	fmt.Println(contains("c", A...))
	fmt.Println(getmax(2, 4, 3, 6, 8))
}
