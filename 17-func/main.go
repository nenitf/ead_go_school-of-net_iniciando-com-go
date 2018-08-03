package main

import "fmt"

func nameR(a string) (c string) {
	c = a

	// return vazio, pois ele retorna var c
	return
}

// retorno dublo
func doubleR(a string, b string) (string, string) {
	return b, a
}

// pode receber diversos valores int de parametro
func variadicF(num ...int) int {
	var res int
	for _, v := range num {
		// fmt.Println(res, v)
		res += v
	}
	return res
}

func main() {
	fmt.Println(nameR("felipe"))

	a, b := doubleR("fel", "ipe")
	fmt.Println(a, b)

	fmt.Println(variadicF(1, 2, 3, 4, 5))
}
