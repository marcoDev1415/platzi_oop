package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() (id int) {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) String() string {

	return fmt.Sprintf("id: %d, Name: %s, Vacation: %t", e.id, e.name, e.vacation)
}

func NewEmployee(id int, name string, vacation bool) *Employee {

	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	e := Employee{}

	fmt.Printf("%v\n", e)

	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}

	fmt.Printf("%v\n", e2)

	e3 := new(Employee)

	fmt.Printf("%v\n", *e3)

	e3.id = 1
	e3.name = "name"

	fmt.Println(e3.String())

	e4 := NewEmployee(4, "Marco", false)

	fmt.Println(e4.String())

}
