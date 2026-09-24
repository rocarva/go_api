package main

import (
	"personalidades/database"
	"personalidades/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
