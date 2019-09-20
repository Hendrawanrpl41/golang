package main

import (
	"fmt"
)

func main(){
	z:=hitung(5,20)
	fmt.Println(z)
}

func hitung(x int, y int) int{
	return x*y
}