package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", Inicio)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}
func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
	//db.StartDbEngine()  //NO FUNCIONA
	//app.StartRoute()

}
