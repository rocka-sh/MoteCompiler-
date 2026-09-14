package analizadorlexico

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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
			break
		}
		if err != nil {
			log.Fatal("Error con el char:", err)
		}
		chars = append(chars, char)
	}

	return chars, nil
}
