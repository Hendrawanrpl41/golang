package main

import (
	"fmt"
	"strconv"
)

func main() {
	nama := "5"
	i, err := strconv.Atoi(nama)

	if err != nil {
		fmt.Println("terjadi error : ",err.Error())
	} else {
		fmt.Println(i)
	}
}