package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	palabras, err := LectorArchivo(args[1])
	if err != nil {
		fmt.Printf("No se pudo abrir el archivo")
	}

	Lexico(palabras)
}
