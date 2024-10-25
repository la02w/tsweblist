package main

import (
	"tsweblist/model"
	"tsweblist/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
