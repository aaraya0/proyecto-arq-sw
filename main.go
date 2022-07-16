package main

import (
	"github.com/aaraya0/arq-software/Integrador1/mvc/app"
	"github.com/aaraya0/arq-software/Integrador1/mvc/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
