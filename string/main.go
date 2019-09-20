package main

import (
	"fmt"
	"strings"
	"strconv"
	
	
)

 func main(){
	var hello = "selamat datang"
	var angka = "100"
	fmt.Println(hello)
	fmt.Println(len(hello))

	// for i := 0;i< len(hello);i++ {
	// 	fmt.Println(hello[i])
	// }
	var helloUpper = strings.ToUpper(hello)
	var helloLower = strings.ToLower(hello)

	fmt.Println(helloUpper)
	fmt.Println(helloLower)
	
	//cek apakah ada dalam variabel atau tidak
	resultContain := strings.Contains(hello, "li")
	fmt.Println(resultContain)

	resultSplit := strings.Split(hello, " ") //dinamyc array
	for _, v := range resultSplit {
		fmt.Println(v)
	}
	
	//conversi string to int
	resultConv, err := strconv.Atoi(angka)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultConv)

 }
