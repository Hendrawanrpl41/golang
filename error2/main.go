package main

import (
	"fmt"
	"errors"
	
)

func main(){
	r, err := div(9, -10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r)
	}

		
}

func div(x,y float64)(float64, error){
	if y < 0 {
		return x/y,  errors.New("jumlah nilai harus positif")
	} else {
		return x/y, nil
	}
}