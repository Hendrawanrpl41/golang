package main

import (
	"fmt"
)

func main(){
	p := &Person {
		id: 1,
		name : "hendrawan",
		age : 21,
	}
	fmt.Println(p.GetId())
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
	p.SetId(3)
	fmt.Println(p.GetId())
	p.SetName("sueng")
	fmt.Println(p.GetName())
	p.SetAge(33)
	fmt.Println(p.GetAge())
}

type Person struct{
	id int
	name string
	age int	
}

func (p *Person)GetId()int{
	return p.id
}
func (p *Person)GetName()string {
	return p.name
}
func (p *Person)GetAge()int {
	return p.age
}

//set
func (p *Person)SetId(newId int){
	p.id = newId
}
func (p *Person)SetName(newName string){
	p.name = newName
}
func (p *Person)SetAge(newAge int){
	p.age = newAge
}