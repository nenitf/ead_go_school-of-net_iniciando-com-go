package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Me diga somente uma palavra!")
	} else {
		fmt.Println(os.Args[1])
	}
}
