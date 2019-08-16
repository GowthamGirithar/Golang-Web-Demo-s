package main

import (
	"context"
	"net/http"
)

func main() {
	http.HandleFunc("/test", testContextValues)
	http.ListenAndServe(":8080", nil)
}

func testContextValues(writer http.ResponseWriter, request *http.Request) {
	// valid for request
	// you will get the new context and from which you can access
	ctx := context.WithValue(request.Context(), "User", "Gowtham")
	println(ctx.Value("User").(string))
	writer.Write([]byte("Test context with values"))

}
