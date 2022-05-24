package main

import (
	//"fmt"
	//	"database/sql"
	"fmt"
	"integrado1/db"
	"log"
	"net/http"
	//_ "github.com/go-sql-driver/mysql"
)

/*func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Password := ""
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Password+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}*/

//var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {

	http.HandleFunc("/", Inicio)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}
func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
	db.StartDbEngine()
	//app.StartRoute()
	//conexionEstablecida := conexionDB()
	//	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO productos(Nombre, Imagen, Stock) VALUES ('Producto 4','IMG0292', 15")
	//if err != nil {
	//	panic(err.Error())
	//}
	//insertarRegistros.Exec()

}
