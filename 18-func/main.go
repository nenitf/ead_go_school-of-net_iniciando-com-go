package main

import "fmt"

// closure: seta o escopo de dentro de uma func

// o retorno é uma função que retorna um inteiro
func funcInsideFunc() func() int {
	a := 15
	return func() int {
		return a
	}
}

func main() {
	x := 10
	fmt.Println(x)

	// func anonima
	add := func() int {
		x += 2
		return x
	}

	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(x)
	fmt.Println("---------------")

	fmt.Println(funcInsideFunc())
}
