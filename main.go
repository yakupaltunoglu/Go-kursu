package main

import (
	"golesson/restful"
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
	// sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)
	// fmt.Println("Toplam: ", sonuc1)
	// fmt.Println("Çıkarım: ", sonuc2)
	// fmt.Println("Çarpım: ", sonuc3)
	// fmt.Println("Bölüm: ", sonuc4)
	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic())
	// sayilar := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	//maps.Demo1()
	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Maindeki sayı:", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki sayı:", sayilar[0])
	// structs.Demo2()

	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)

	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	// carpim := ciftSayiToplam * tekSayiToplam

	// fmt.Println("Çarpım :", carpim)

	// time.Sleep(5 * time.Second)

	// fmt.Println("Main bitti")

	// defer_statement.Demo3()
	// error_handling.Demo1()

	// fmt.Println(error_handling.TahminEt2(102))
	// string_func.Demo1()
	// string_func.Demo2()
	restful.Demo2()

}
