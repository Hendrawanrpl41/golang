package main
import (
	"fmt"
)

func main(){

	var players []Player
	players = []Player{Player{id: 1, name : "hendrawan"}, Player{id : 2, name : "sueng"}}
	players = append(players, Player{id:3, name : "bonyok"})
	for _, v := range players{
		fmt.Println(v)
	}
}


type Player struct{
	id int
	name string
}