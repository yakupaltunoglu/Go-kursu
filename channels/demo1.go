package channels

import (
	"fmt"
	"time"
)

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("çift sayı toplama çaılışıyor")
		time.Sleep(1 * time.Second)
	}
	ciftSayiCn <- toplam
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("tek sayı toplama çaılışıyor")
		time.Sleep(1 * time.Second)
	}
	tekSayiCn <- toplam
}
