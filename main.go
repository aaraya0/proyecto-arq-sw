package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//db.RunDB()
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/home", Productos)
	http.HandleFunc("/cart", Carrito)
	http.HandleFunc("/orden", Orden)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)

}
func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")

}
func Productos(w http.ResponseWriter, r *http.Request) {
	//	controller.ShowProd(w)
	//fmt.Fprintf(w, "Hola mundo2")
}

func Carrito(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hola mundo3")
}
func Orden(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hola mundo4")
}
