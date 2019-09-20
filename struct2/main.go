package main

import(
	"fmt"
)
func main() {
	p := Person{
		id:1,
		name : "hendrawan",
		age : 21,
	}	
	printPerson(p)
}

type Person struct{
	id int
	name string
	age int
}

func printPerson(p Person){
	fmt.Println(p.id)
	fmt.Println(p.name)
	fmt.Println(p.age)
}
