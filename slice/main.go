package main

import (
	"fmt"
)

func main(){
	mySlice := []int {10,11,12,13,14,15}
	mySliceStr := []string {"hendrawan","sueng","maneh"}
	
	fmt.Println(mySlice[0])
	for i, v := range mySlice {
		fmt.Println(i, v)
	}
	

	mySliceStr = append(mySliceStr, "ratih") //ADD DATA
	for _,v := range mySliceStr{
		fmt.Println(v)
	}

}