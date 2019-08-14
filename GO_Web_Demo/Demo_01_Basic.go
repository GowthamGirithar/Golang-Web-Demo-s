package main

import (
	"fmt"
	"net/http"
)
//when you call from the browser , sometimes favicon.io requests also made
//so handler can be called twice- try using curl or postman for this problem
func main() {
	http.HandleFunc("/status", getStatus)
	http.HandleFunc("/", home)
	http.ListenAndServe("127.0.0.1:8081", nil)  //ip and port or :8081 is enough

}
func getStatus(writer http.ResponseWriter, request *http.Request){
	println("called")
	fmt.Fprintf(writer, "Responses")
}

func home(writer http.ResponseWriter, request *http.Request){
	println("home called")
	fmt.Fprintf(writer, "Welcome to Home Page")
}


//why we should not use the DefaultServeMux?
//it is global one and anyone can define the route
//tomorrow you are using some third party dependency and that has the init function which defines some route
//it is a problem. so we have to use the our NewServeMux which is in next exercise

//var defaultServeMux ServeMux
//var DefaultServeMux = &defaultServeMux
