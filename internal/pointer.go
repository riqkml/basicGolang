package internal

import "log"

type Address struct {
	City     string
	Province string
	Country  string
}

func ChangeCityToJakarta(address *Address) {
	address.City = "Jakarta"
}

func Pointer() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address1.City = "Purwakarta"

	address3 := new(Address)
	address3 = &address1
	address3.City = "Malang"
	address3.Province = "Jawa Timur"
	address3.Country = "Indonesia"

	log.Println("address1 ", address1)
	log.Println("address2 ", address2)
	log.Println("address3 ", address3)
}
