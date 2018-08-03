package main

import "fmt"

func main() {
	nome := "felipe"

	switch nome {
	case "diogo":
		fmt.Println("olá diogo")
	case "felipe":
		fmt.Println("hey eu msm")
	default:
		fmt.Println("ué, qm é vc")
	}
}
