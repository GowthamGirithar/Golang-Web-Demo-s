package main

import "net/http"

// it is like chain of filter before executing and after
// we can use this concept for logging

//middleware function should follow (next http.Handler) http.Handler because we may use chain of handler and return the handler
func main() {
	mux := http.NewServeMux()
	homeContent := http.HandlerFunc(homeContent) // its type of handler func .. which will be converted into handler
	mux.Handle("/", middlewareOne(middlewareTwo(homeContent)))
	http.ListenAndServe(":8080" , mux)
}

func middlewareTwo(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			println("Begining of middlewareTwo")
			if r.URL.Path == "/favicon.ico"{
				println("favicon request will not be proceeded")    // when you hit the url from browser , it send favicon.ico also as / so for these request im not proceeding
																			//best place for scope check too
				return
			}
			next.ServeHTTP(w, r)
			println("ending of middlewareTwo")
	})
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		println("Begining of middlewareOne")
		next.ServeHTTP(w, r)
		println("ending of middlewareOne")
	})
}

func homeContent(w http.ResponseWriter, r *http.Request){
	println("Final One")
	w.Write([]byte("Home Page"))
}

/***
Output
Begining of middlewareOne
Begining of middlewareTwo
Final One
ending of middlewareTwo
ending of middlewareOne

 */