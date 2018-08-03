package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("---------------------")
	// white
	x := 0
	for x <= 10 {
		fmt.Println(x)
		x++
	}

	// laço infinito
	y := 0
	// não precisa ser "for true"
	for {
		fmt.Println(y)
		y++
		if y == 100 {
			// interrompe o for
			break
		}
	}
}
