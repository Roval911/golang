package main

import "fmt"

func main() {
	workweek := []string{"пн", "вт", "ср", "чт", "пт"}
	restday := []string{"сб", "вс"}
	workweek = append(workweek, restday...)
	fmt.Println(workweek)

}
