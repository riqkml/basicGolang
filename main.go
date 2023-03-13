package main

import (
	"examplego/internal"
	"log"
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

	log.Println("city ", riqki.Name)
}
