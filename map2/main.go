package main

import (
	"fmt"
)

func main(){
	mapPlayer := make(map[int]Player)
	mapPlayer[1] = Player{id : 1, name : "hendrawan"}
	mapPlayer[2] = Player{id : 2, name : "jupe"}

	for _, v := range mapPlayer{
		fmt.Println(v.name)
	}
}
type Player struct{
	id int
	name string
}