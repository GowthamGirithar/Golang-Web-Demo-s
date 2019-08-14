package main

import (
	"net/http"
	"time"
)

/**
any handler should have this method or any function which this arguement can be converted into the handler
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
 */


//handler function with arguement of (writer http.ResponseWriter, request *http.Request) can be converted into hander type
//method which returns the handler
//handler function have implemented the handler interface
func homePageContent(welcomeStr string) http.Handler{
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(welcomeStr))
	})
}
//normal handler function
func testPage(writer http.ResponseWriter, request *http.Request){
	writer.Write([]byte("test page"))
}

type timeFormat struct{
	Format string
}
//here time format type implemented the handler method so it became type of handler
func (tf *timeFormat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The time is: " + time.Now().Format(tf.Format)))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/home" ,homePageContent("Hello how are you"))  // you are getting handler
	mux.HandleFunc("/test" ,testPage) //u are calling the handler function
	timeFormat := &timeFormat{Format:time.RFC850}  //here you are using your own type and which implement the handler method so you can use in handle
	mux.Handle("/time",timeFormat)
	redirectHandler := http.RedirectHandler("http://www.google.com",307)
	mux.Handle("/google" , redirectHandler)   // default handler usage
	http.ListenAndServe(":8080" , mux)
}


//Learnings of handler:
//1.handler should have the handler interface implementation
//2.whenever any function which have the same method arguements like the handler ,it can be used in handler
//3.if you want to pass the data , you can use the handler which should return the handler - return will be of type handler
//4.if we have type which implements the handler function, then we can use
//5. we can also use the handlers provided by the golang like RedirectHandler, error handler etc.,



