package main

import "fmt"

var e1 string = "escopo de pacote"

func main() {
	fmt.Println(e1)
	funcEscopo()

	//Para o go encontrar a var e3
	//go run *.go
	fmt.Println(e3)

	//Ou usar "go build" para gerar binario de mesmo nome da pasta, e executa-lo
	//"go clean" apaga o binario
}

func funcEscopo() {
	e2 := "escopo de func"
	fmt.Println(e2)
}
