package main

import (
	"fmt"
	"net/http"
)

//demo with NewServeMux
// advantage of using the NewServeMux
func main() {

	mux := http.NewServeMux()  //if we dont define default serve mux is used
	mux.HandleFunc("/status", health)
	mux.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", mux) //-> HTTP1.1

	//to run in http2
	//http.ListenAndServeTLS(":8080", "server.crt", "server.key", mux)

}

func health(writer http.ResponseWriter, request *http.Request) {
	println("called")
	fmt.Fprintf(writer, "UP")
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	println("called")
	fmt.Fprintf(writer, "UP")
}

//why we should use the NewServeMux?
//DefaultServeMux is global one and anyone can define the route
//tomorrow you are using some third party dependency and that has the init function which defines some route
//it is a problem. so we have to use the our NewServeMux which is in next exercise
