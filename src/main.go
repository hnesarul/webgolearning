package main

import "github.com/gorilla/mux"

func main() {

	r := mux.NewRouter()

	r.HanleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)

}
