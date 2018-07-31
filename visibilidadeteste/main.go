// este pacote será usado para testar a pasta 10-visibilidade
// import "visibilidadeteste"

// o pacote deve estar "buildado" e na raiz do go/src, $GOPATH
// ESSA PASTA AQUI ONDE ESTÁ NÃO FUNCIONA, PRECISA ESTAR NA RAIZ DE GO/SRC
package visibilidadeteste

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// Só é possivel acessar funções que iniciem com letra maiuscula de pacotes externos
// fmt.Println("tipo esse")
func printUm() {
	fmt.Println("Vc não tem acesso a essa func por um pacote externo")
}

func PrintDois() {
	fmt.Println("Vc conseguiu me acessar hohoho")
}
