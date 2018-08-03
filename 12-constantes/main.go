package main

import "fmt"

// constantes com letra maiusculua para serem publicas
const ctn1 int = 10
const (
	ctn2 int    = 12
	ctn3 string = "felipe"
)
const ctn4 = 14.5

func main() {
	fmt.Println(ctn1)
	fmt.Printf("%T \n", ctn2)
	fmt.Println(ctn3)
	fmt.Println(ctn4)
}
