package main

import "fmt"

func main() {
	var a int = 1
	f := float64(a)
	var b float64 = 1.23
	var c bool = true
	var d string = "vin"
	const e = "msn"

	fmt.Printf("Variable %v is type of %T\n", a, a)
	fmt.Printf("Variable %v is type of %T\n", b, b)
	fmt.Printf("Variable %t is type of %T\n", c, c)
	fmt.Printf("Variable %v is type of %T\n", d, d)
	fmt.Printf("Variable %v is type of %T\n", e, e)
	fmt.Printf("Variable %f is type of %T\n", f, f)

}
