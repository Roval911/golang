package main

import (
	"fmt"
	"strconv"
)

func main() {
	first := "104"
	second := 35

	s := strconv.Itoa(second)
	f, _ := strconv.Atoi(first)
	fmt.Println(s, f)
}
