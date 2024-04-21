package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", handleRequest)

	log.Println("Escuchando en 127.0.0.1:8080...")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Definir la ruta base del directorio de PHP
	baseDir := "C:\\golang\\server\\php"

	// Si la ruta solicitada es la raíz ("/"), servir index.php
	if r.URL.Path == "/" {
		scriptPath := baseDir + "\\index.php"
		serveScript(w, r, scriptPath)
		return
	}

	// Para otras rutas, servir el archivo solicitado
	scriptPath := baseDir + r.URL.Path
	serveScript(w, r, scriptPath)
}

func serveScript(w http.ResponseWriter, _ *http.Request, scriptPath string) {
	// Ejecutar el script PHP
	out, err := exec.Command("C:\\php\\php-cgi.exe", "-f", scriptPath).Output()
	if err != nil {
		log.Printf("Error al ejecutar el script PHP: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Escribir la salida del script PHP en la respuesta HTTP
	fmt.Fprintf(w, "%s", out)

	// Imprimir la salida del script PHP para depuración
	log.Printf("Salida del script PHP: %s", out)
}
