package main

import "fmt"

func main() {
	week := []string{"пн", "вт", "ср", "чт", "пт", "сб", "вс"}
	workdays := make([]string, 10, 10)
	workdays = week[0:5]
	week = week[5:7]

	fmt.Println(workdays)
	fmt.Println(week)

}
