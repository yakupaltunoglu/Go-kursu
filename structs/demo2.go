package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func Demo2() {
	c := customer{firstName: "Engin", lastName: "Demiroğ", age: 35}
	c.save()
	c.update()
}

func (c customer) save() {
	fmt.Println("eklendi", c.firstName)
}
func (c customer) update() {
	fmt.Println("güncellendi", c.firstName)
}
