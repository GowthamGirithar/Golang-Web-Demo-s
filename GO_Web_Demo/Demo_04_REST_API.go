package main

import (
	"encoding/json"
	"net/http"
)

type EmployeeDto struct {
	EmployeeId string
	Name       string
}

//Demo for posting and working on request objects
//use decode over unmarshal if you are reading from the request
//unmarshal if you have full json already in variable
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/Employee", handleEmployee)
	mux.HandleFunc("/Employee/{id}", handleEmployee)
	http.ListenAndServe(":8080", mux)
}



func handleEmployee(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
	case "POST":
			emp := EmployeeDto{}
			errs := json.NewDecoder(request.Body).Decode(&emp)// interface always require the address type
			checkError(errs)
			println("the body is " ,emp.Name)
	case "PUT":
			pathValues := request.URL.Query()  //return map
			empid := pathValues.Get("id")
			println("the emp id is ", empid)

	case "DELETE":
	default:
		println("None of the methods other than above are allowed")
	}
}

//Use json.Decoder if your data is coming from an io.Reader stream, or you need to decode multiple values from a stream of data.
//Use json.Unmarshal if you already have the JSON data in memory.
/**
example to use the unmarshal
jsonBody, err :=ioutil.ReadAll(request.Body)
			defer request.Body.Close()
			checkError(err)
			emp := EmployeeDto{}
			errs := json.Unmarshal([]byte(jsonBody), &emp)  // interface always require the address type
			checkError(errs)
			println("the body is " ,emp.Name)
 */

func checkError(err error){
	if err != nil{
		println("the error is ", err.Error())
	}
}
