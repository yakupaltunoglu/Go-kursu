package for_range

import "fmt"

//1-10 arasındaki sayılardan tek olanları toplayan program
//forrange
func Demo2() {
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for i, sayi := range sayilar {
		if i%2 != 0 {
			toplam = toplam + sayi
		}
	}
	fmt.Println("Toplam:", toplam)
}
