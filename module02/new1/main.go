package main

import "fmt"

func main() {
	var AmericanVelocity float64
	var EuropeanVelocity float64
	EuropeanVelocity = 120.4 / 1000 * 3600
	AmericanVelocity = 130.0 * 2.237
	fmt.Println(AmericanVelocity)
	fmt.Println(EuropeanVelocity)
}
