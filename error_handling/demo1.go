package error_handling

import (
	"fmt"
	"os"
)

// type assertion
func Demo1() {
	f, err := os.Open("demo11.txt")
	//nil
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamad─▒", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamad─▒")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
