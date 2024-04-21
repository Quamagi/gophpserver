package main

import (
	"fmt"
	"io"
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
	// Configurar CORS
	addCORSHeaders(w, r)

	// Manejar solicitudes de opciones CORS
	if r.Method == http.MethodOptions {
		return
	}

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

func addCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func serveScript(w http.ResponseWriter, r *http.Request, scriptPath string) {
	// Ejecutar el script PHP
	cmd := exec.Command("C:\\php\\php-cgi.exe", "-f", scriptPath)

	// Pasar el cuerpo de la solicitud al script PHP para métodos POST, PUT y DELETE
	if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cmd.Stdin = &Pipe{bytes: body}
	}

	out, err := cmd.Output()
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

// Pipe es un tipo auxiliar que implementa la interfaz io.Reader
type Pipe struct {
	bytes []byte
	pos   int
}

// Read implementa el método Read de io.Reader
func (p *Pipe) Read(b []byte) (int, error) {
	n := copy(b, p.bytes[p.pos:])
	p.pos += n
	if n == 0 {
		return 0, io.EOF
	}
	return n, nil
}
