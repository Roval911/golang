package main

func main() {
	var A *int
	B := 12
	A = &B
	println(*A)
	B = 14
	println(B)

}
