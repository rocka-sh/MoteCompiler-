package analizadorlexico

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

// abre el archivo y lo hace un arreglo de caracteres, agregando %s \t \n
func LectorArchivo(p string) ([]rune, error) {
	rutaAbsoluta, err := filepath.Abs(p)
	file, err := os.Open(rutaAbsoluta)

	if err != nil {
		fmt.Println("no se pudo abrir el archivo", err)
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var chars []rune

	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			fmt.Println("fin archivo, EOF")
			break
		}
		if err != nil {
			log.Fatal("Error con el char:", err)
		}
		chars = append(chars, char)
	}

	return chars, nil
}

// separa por palabras
func SepararPorEspacios(caracteres []rune) [][]rune {
	var palabras [][]rune
	var palabraActual []rune

	for _, char := range caracteres {
		if unicode.IsSpace(char) {
			if len(palabraActual) > 0 {
				palabras = append(palabras, palabraActual)
				palabraActual = nil
			}
		} else {
			palabraActual = append(palabraActual, char)
		}
	}

	if len(palabraActual) > 0 {
		palabras = append(palabras, palabraActual)
	}

	return palabras
}
