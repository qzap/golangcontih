package main

import(

	"goprogram/golangcontoh/db"
	"goprogram/golangcontoh/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
