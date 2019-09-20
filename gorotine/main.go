package main

import (
	"fmt"
)
//mirip dengan promise untuk mencegah async

func main() {
	done := make(chan bool)
	go hello(done)
	<- done
}

func hello(done chan bool){
	fmt.Println("hello go")
	done <- true
}

