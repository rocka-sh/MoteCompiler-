package analizadorlexico

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func abrirArchivo(p string) (*os.File, error) {
	rutaAbsoluta, err := filepath.Abs(p)
	file, err := os.Open(rutaAbsoluta)

	if err != nil {
		fmt.Println("no se pudo abrir el archivo", err)
		return nil, err
	}
	return file, nil
}

// TODO: hacer que retornen los chars
func lectorArchivo(f *os.File) []rune {
	reader := bufio.NewReader(f)
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

	return chars
}
