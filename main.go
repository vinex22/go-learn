package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 1
	f := float64(a)
	var b float64 = 1.23344
	var c bool = true
	var d string = "Aseem"
	const e = "Gaur ji Bandook wale"
	g := (d == e)
	h := d + " " + e
	fmt.Printf("Variable a = %v is type of %T\n", a, a)
	fmt.Printf("Variable b = %v is type of %T\n", b, b)
	fmt.Printf("Variable c = %t is type of %T\n", c, c)
	fmt.Printf("Variable d = %v is type of %T\n", d, d)
	fmt.Printf("Variable e = %v is type of %T\n", e, e)
	fmt.Printf("Variable f = %f is type of %T\n", f, f)
	fmt.Printf("Variable g  = %v is type of %T\n", g, g)
	fmt.Printf("Concationation of d and e is %v\n ", h)
	fmt.Print(ToMorseCode(h))
	var x, y float64 = 26.22, 26.22
	var z float64 = math.Mod(x, y)
	fmt.Print(z)
	if x != y {
		fmt.Printf("not equal")
	} else {
		fmt.Printf(" equal")
	}
	var fruit [5]string
	grades := [...]int{1, 2, 3}

	fmt.Print(len(fruit))
	fmt.Print(len(grades))
	fmt.Println(grades[1])
}
