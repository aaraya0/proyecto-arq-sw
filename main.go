package main

import (
	"github.com/aaraya0/arq-software/Integrador1/app"
	"github.com/aaraya0/arq-software/Integrador1/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
