package conditionals

import (
	"fmt"
)

func Demo3() {
	//Üç adet int değişken tanımlayınız.
	//Ekrana en büyük olanı yazdırınız.
	var sayi1, sayi2, sayi3 int = 10, 5, 18
	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}

	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En büyük sayı :%v ", enBuyuk)

}
