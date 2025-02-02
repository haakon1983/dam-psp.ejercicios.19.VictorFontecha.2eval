package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// Crear una solicitud GET
	resp, err := http.Get("http://localhost:8025/assets/img/logo.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	//Crear un lector para leer línea a línea
	reader := bufio.NewReader(resp.Body)
	headerDone := false

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		//Mostrar la cabeceras cabeceras HTTP
		if !headerDone {
			if strings.TrimSpace(line) == "" {
				headerDone = true
			} else {
				fmt.Print("CABECERA>> %s", line)
			}
		} else {
			//Detectar el inicio del contenido binario de la imagen PNG
			if strings.HasPrefix(line, "\x89PNG") {
				fmt.Println("BINARIO>> %s", line)
			}
		}
	}
}
