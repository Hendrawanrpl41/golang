package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	// h := http.NewServeMux()
	// http.Get("localhost:3000/users")
	http.HandleFunc("/", hello)
	http.HandleFunc("/users", penutup)
	http.ListenAndServe(":3000",nil)
	log.Println(nil)

}

func hello(res http.ResponseWriter, req *http.Request ){
	fmt.Fprintf(res, "hello selamat datang")
}

func penutup(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, req.Form)
}