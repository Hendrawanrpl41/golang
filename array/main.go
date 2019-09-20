package main

import (
	"fmt"
)

func main(){
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	arrStr := [3]string {"hendrawan", "jupe","mufiq"}

	arrMulti := [3][3]int {{1,2,3}, {4,5,6},{7,8,9}}
 
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	} 
	fmt.Println(arrStr[2])
	fmt.Println(arrMulti[2][2])
}