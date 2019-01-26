package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	yuki := person{
		firstName: "Yuki",
		lastName:  "Lee",
		contactInfo: contactInfo{
			email: "yuki@gmail.com",
			zip:   1000,
		},
	}

	yukiPointer := &yuki
	yukiPointer.updateFirstName("Nani")
	yuki.print()
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
