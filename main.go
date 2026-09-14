package main

import (
	adf "MoteCompiler/analizadorlexico"
	pr "MoteCompiler/analizadorlexico/palabrasreservadas"
	"fmt"
	"os"
)

func main() {
	argumentos := os.Args
	archivo, err := adf.LectorArchivo(argumentos[1])
	if err != nil {
		fmt.Println("No se pudo abrir el archivo:", err)
		return
	}

	palabras := adf.SepararPorEspacios(archivo)

	for _, palabra := range palabras {
		pr.PalabrasReservadas.Tokenizador(palabra)
	}
}
