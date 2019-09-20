package main

import "fmt"

func main(){

	f := func(v string)bool{
		return v=="golang"
	}
	resutl := match("golang", f)
	fmt.Println(resutl)
}
//membuat function yang mengembalikan nilai boolean
func match(v string, f func(string)bool)bool{
	if f(v){
		return true
	}
	return false
}