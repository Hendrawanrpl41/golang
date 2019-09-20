package main

import(
	"fmt"
)
func main(){
	
	var mapPerson map[int]string
	mapPerson = make(map[int]string)
	mapPerson[1] = "alex"
	mapPerson[2] = "curung"
	mapPerson[4] = "dwik"

	for k, v := range mapPerson{
		fmt.Println(k, v)
	}

	cek, err := mapPerson[0]
	
	if !err {
		fmt.Println("data tidak ditemukan")
	}else{
		fmt.Println("data di temukan")
		fmt.Println(cek)
	}
}