package main

import (
	"examplego/internal"
	"log"
	"os"
)

func main() {
	//pointer
	internal.Pointer()
	home := internal.Address{}

	// pointer function
	internal.ChangeCityToJakarta(&home)

	// pointer method
	riqki := internal.Man{Name: "Riqki"}
	riqki.GetMarried()

	// package init
	con := internal.GetDatabase()

	// package OS
	_ = os.Args
	name, _ := os.Hostname()

	log.Println("db con = ", con)

	id := os.Getenv("IDAPP")
	log.Println(name, id)

	internal.GenerateList()
}
