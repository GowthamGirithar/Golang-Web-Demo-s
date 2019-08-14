package main

import (
	"net/http"
	"net/http/httptest"
	"time"
)
//dont use the default client in production because which dont have the timeout
//so it will wait till it get the response
//we should have the timeout for any request
// default client is var DefaultClient = &Client{}

//let us consider one go routine is executing the get call and service is take too much time and after few attempts
// our application will resources gets reduced and server will down - (considering many ppl request for the same functionality)


//http packages use http1.1 for client
//to use HTTP2.0 use "golang.org/x/net/http2"
func main() {

	//server for testing purpose - here server will respond after 1 hour  - this is for demo purpose
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Hour) //respond after hours
	}))
	println(svr.URL)
	defer svr.Close()


	//my client will will only wait for 10 seconds and throw timeout error
	httpClient := http.Client{
			Timeout:       10 *time.Second,   //this is 0 means no timeout in default one
	}
	res , err :=httpClient.Get(svr.URL)  //server url
	if err != nil{
		println("the error is " ,err.Error())
	}
	println("the response is  " , res.Status)

	//http.Post("","application/json",json.Decoder{}.Buffered())
	//jsonValue, _ := json.Marshal(values)
	//resp, err := http.Post(authAuthenticatorUrl, "application/json", bytes.NewBuffer(jsonValue))


}
