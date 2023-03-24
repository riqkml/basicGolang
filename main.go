package main

import (
	"examplego/internal"
	"fmt"
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

	sample := internal.Sample{Name: "Riqki", Age: 23}
	fmt.Println("sample ", sample.IsValid())
}
