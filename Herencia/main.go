package main

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {

	ftmEmploye := FullTimeEmployee{}
	ftmEmploye.name = "Marco"
	ftmEmploye.age = 27
	ftmEmploye.id = 1

}
