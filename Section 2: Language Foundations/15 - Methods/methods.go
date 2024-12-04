package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving user %s on database\n", u.name)
}

func (u user) isOfLegalAge() bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
}

func main() {
	u := user{"John", 20}
	fmt.Println(u)
	u.save()

	isOfLegalAge := u.isOfLegalAge()
	fmt.Println(isOfLegalAge)

	u.birthday()
	fmt.Println(u.age)
}
