pakage headler


//Home

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome to the Home!")
}

//Saludo

func SaludoHandler(w http.ResponseWriter, r /http.Request) {

	nombre := r.URL.Query().Get("nombre")
	if nombre ==""{
		nombre = "extraño"
}
	fmt.Fprintf(w, "Hola, %s! ",  nombre)
}

//About
func AboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Esta API esta en Go Puro ")
}