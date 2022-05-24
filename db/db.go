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
	insert, err := db.Query("INSERT INTO `test` (`Nombre`) VALUES ('Test2')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Yay, values added!")
}
