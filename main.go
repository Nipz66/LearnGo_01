package main

import (
	"fmt"
)

const s string = "Constance"

func main() {

	fmt.Println(s)

	const n = 5000

	const d = 300000 / n

	fmt.Println(d)

	//var use create one or more variables

	var a = "Nipun Anupama"
	fmt.Println(a)

	// You can declare multiple variables

	var c, b int = 1, 2
	fmt.Println(c, b)

	//type of initialized variables

	var x = true
	fmt.Println(x)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable

	// n := "Apple "
	// fmt.Println(n)

}
