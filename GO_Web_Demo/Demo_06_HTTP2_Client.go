package main

import (
	"golang.org/x/net/http2"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {


		//server for testing purpose - here server will respond after 1 hour  - this is for demo purpose
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Hour) //respond after hours
		}))
		println(svr.URL)
		defer svr.Close()


		//my client will will only wait for 10 seconds and throw timeout error
		httpClient := http.Client{
			Transport: &http2.Transport{},
			Timeout:       10 *time.Second,   //this is 0 means no timeout in default one
		}
		res , err :=httpClient.Get(svr.URL)  //server url
		if err != nil{
			println("the error is " ,err.Error())
		}
		println("the response is  " , res.Status)

}

/**
Note:
1. http.RoundTripper is default used and it is used to cache and use of it
2. httptrace is used to trace the following events
	Connection creation
	Connection reuse
	DNS lookups
	Writing the request to the wire
	Reading the response
3. Both 1 and 2 are used internally by http client
4. By default timeout is 0 for waiting for response. so , please use the timeout otherwise once connection is established no timeout and it will wait till it get response



*/
