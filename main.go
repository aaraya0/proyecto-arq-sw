package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaraya0/arq-software/Integrador1/db"
)

func main() {

	http.HandleFunc("/", Inicio)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}
func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")

	db.RunDB()

}
