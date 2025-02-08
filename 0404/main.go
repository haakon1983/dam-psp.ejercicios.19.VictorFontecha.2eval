package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Crear una solicitud GET
	resp, err := http.Get("http://localhost:8025/assets/img/logo.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Crear un buffer temporal para almacenar el contenido binario
	var imagen bytes.Buffer
	_, err = io.Copy(&imagen, resp.Body)
	if err != nil {
		fmt.Println("Error al copiar los datos:", err)
		return
	}

	// Crear un archivo de salida
	archivo, err := os.Create("salida.png")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close()

	// Escribir el contenido del buffer en el archivo
	_, err = imagen.WriteTo(archivo)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Imagen guardada como salida.png")
}
