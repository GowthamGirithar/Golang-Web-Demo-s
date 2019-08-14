package main

import (
	"log"
	"net/http"
)

//TO USE HTTP 2.0 we have to use with authentication
//only http client will require the extra http2 dependency to convert into the http 2 call
//golang provides http2 for ListenAndServeTLS and not for ListenAndServe
//example of server push
func main() {
	mux := http.NewServeMux()
	print("test")
	mux.HandleFunc("/", testHTTP2)   //when asking / , send all the thing

	//generate the certificate using https://www.samltool.com/self_signed_certs.php
	//used sha256
	serverError := http.ListenAndServeTLS("127.0.0.1:8080" ,"server.cert","server.key",mux)
	if serverError!= nil{
		println("the error is  " ,serverError.Error())
	}
}

func testHTTP2(w http.ResponseWriter, r *http.Request){
	println("http2 REQUEST" , r.URL.Path)

	//server push:
	if r.URL.Path == "/test2"{
		println("test2 end point")
		w.Write([]byte("Responses from 2"))
		return  //very important otherwise recursive happens
		//this example is like sending index.html , send app.css or other css and static files
	}
	// Server push must be before response body is being written.
	// In order to check if the connection supports push, we should use
	// a type-assertion on the response writer.
	// If the connection does not support server push, or that the push fails we
	// just ignore it - server pushes are only here to improve the performance for HTTP2 clients.
	pusher , ok := w.(http.Pusher)
	if ok{
		err := pusher.Push("/test2", nil)
		if err != nil {
			log.Printf("Failed push: %v", err)
		}
	}

	print("test1 end point")
	w.Write([]byte("Responses from 1"))
}