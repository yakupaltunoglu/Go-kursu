package structs

import "fmt"

func Demo1() {
	// fmt.Println(product{"Bilgisayar", 500, "xyz"})
	// fmt.Println(product{"Bilgisayar", 500, ""})
	fmt.Println(product{name: "Bilgisayar"})
}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}
