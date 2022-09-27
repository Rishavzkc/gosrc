package main

import (
	"fmt"
	"log"
)

func main() {

	var myPointerVar *int
	fmt.Println(&myPointerVar)

	cities := make(map[string]string)

	cities["Delhi"] = "Del"
	log.Println(cities)

	EUcountries := []string{"Austria", "Belgium", "Bulgaria"}
	addCountries(&EUcountries)
	log.Println(EUcountries)
}

func addCountries(countries *[]string) {
	*countries = append(*countries, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)

}
