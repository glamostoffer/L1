package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) Speak() {
	fmt.Printf("Hi! My name is %s! I'm %d yo.\n", h.name, h.age)
}

type Action struct {
	Human
	profession string
	salary     int
}

func (a *Action) GetSalary() {
	fmt.Printf("My name is %s and I'm an %s. My salary is %d. \n", a.name, a.profession, a.salary)
}

func main() {
	a := Action{
		Human: Human{
			name: "Mike",
			age:  32,
		},
		profession: "engineer",
		salary:     4600,
	}

	a.Speak()
	a.GetSalary()
}
