package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Driver := "mysql"
	Usuario := "root"
	Password := ""
	Nombre := "sistema"
	db, err := sql.Open(Driver, Usuario+":"+Password+"@tcp(127.0.0.1)/"+Nombre)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	leer, err := db.Query(`SELECT * FROM productos WHERE ID=4`)
	if err != nil {
		panic(err.Error())
	}
	defer leer.Close()
	fmt.Println("successsssss")
	fmt.Println(leer)
}
