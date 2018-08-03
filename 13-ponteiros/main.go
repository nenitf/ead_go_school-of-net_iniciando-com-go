package main

import "fmt"

func semPointer(a int) int {
	a += 100
	return a
}

// recebe endereço de memoria inteiro
func comPointer(a *int) int {
	// pega valor do endereço de memoria e aumenta em 100
	*a += 100
	return *a
}

func main() {
	// endereço de memoria com  & antes da variavel
	x := 10
	fmt.Println(&x)

	// copiar o endereçamento de x
	y := &x
	fmt.Println(y)

	// * pega o valor que esta na variavel e o considera como endereço; depois imprime o valor que está no endereço referenciado
	// semelhante ao $$var do php, porem com ponteiro
	fmt.Println(*y)

	// muda o valor que esta apontado para 20
	*y = 20
	fmt.Println(y)
	fmt.Println(x)

	fmt.Println("-----------------------")
	// recebe endereçamento de memória que só recebe inteiro
	var z *int = &x
	// imprime endereço da memoria
	fmt.Println(z)
	// imprime valor da memoria
	fmt.Println(*z)

	fmt.Println("-----------------------")
	b := 12

	fmt.Println(semPointer(b))
	fmt.Println(b)

	fmt.Println(comPointer(&b))
	fmt.Println(b)
}
