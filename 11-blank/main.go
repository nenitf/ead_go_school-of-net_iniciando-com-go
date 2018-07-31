package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/*
		uma função pode retornar dois valores, o resultado e o erro
		quando não foi tratado o erro com if err != nil, a var err vai ser citada como não utilizada

		para evitar tal situação é comum deixar o nome da var de _
		pois o go não gera erro ao compilar uma var chamada _ que não foi usada
	*/

	// gera erro no err
	//res, err := http.Get("http://google.com")
	res, _ := http.Get("http://google.com")
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", body)
}
