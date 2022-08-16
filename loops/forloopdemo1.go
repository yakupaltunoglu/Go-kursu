package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya!"

	for i := 1; i < 5; i++ {
		fmt.Println(metin)
	}
	fmt.Println("Bitti")
}
