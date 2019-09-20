package main
import (
	"fmt"
)

type Vector struct{
	x int
	y int

}
type Player struct{
	id int
	nama string

}

func main(){
	var v Vector 
	v.x = 10
	v.y = 5
	fmt.Println(v)

	player1 := Player{id : 1,nama : "Hendrawan"}
	fmt.Println(player1)

}