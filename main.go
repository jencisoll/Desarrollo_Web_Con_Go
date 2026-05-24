package main

import (
	"fmt"
	"net/http"
)

func main() {

	//Cuando llega un requesta /, llama a esta función
	http.HandleFunc("/", homeHandler)
	//Arranca el servidor en el puerto 8080
	//El segundo argumento nil significa "Usa el router por defecto
	fmt.Println("Listening on port 8080,  server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// Un handler siempre recibe estos dos parametros:
// w -> ResponceWriter: para ESCRIBIR la respuesta al cliente
// r ->  Request: contiene TODO lo del request (método, body, headers, etc.)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}
