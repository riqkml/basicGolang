package main

import (
	"examplego/internal"
	"log"
)

func main() {
	internal.Pointer()
	home := internal.Address{}

	internal.ChangeCityToJakarta(&home)

	log.Println("city ", home.City)
}
