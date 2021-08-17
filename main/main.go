package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to Go!! You are in the Home Page!!")
	fmt.Println("Endpoint Hit: homePage")

}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Deren"
	fmt.Fprintf(response, "My name is "+who)
	fmt.Println("Endpoint Hit: ", who)

}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
