package forrange

import "fmt"

func Demo3() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}
	for k, v := range sozluk {
		fmt.Print(k)
		fmt.Println(v)
	}
}
