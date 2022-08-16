package variables

import "fmt"

func Demo1() {
	var metin string = "Merhaba"
	fmt.Println(metin)
	fmt.Println(metin)

	var kdv int = 15
	fmt.Println(kdv)
	fmt.Println(kdv)

	var kdv2 float64 = 5.5
	fmt.Print(kdv2)

	kdv3 := 25
	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T\n", kdv3)

	var durum bool

	var metin1 string = "Engin"
	var metin2 string = "Ahmet"

	durum = metin1 == metin2
	fmt.Println(durum)
}
