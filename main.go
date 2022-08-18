package main

import (
	"fmt"
	"golesson/functions"
)

func main() {
	// variables.Demo1()
	// fmt.Print()
	// conditionals.Demo1()

	// conditionals.Demo3()
	// loops.Demo1()
	// slices.Demo2()
	// functions.SelamVer("engin")
	// var sonuc = functions.Topla(2, 6)
	// fmt.Println(sonuc)
	sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)
	fmt.Println("Toplam: ", sonuc1)
	fmt.Println("Çıkarım: ", sonuc2)
	fmt.Println("Çarpım: ", sonuc3)
	fmt.Println("Bölüm: ", sonuc4)

}
