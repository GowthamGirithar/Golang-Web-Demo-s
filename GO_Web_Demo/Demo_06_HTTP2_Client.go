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
