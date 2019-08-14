package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

//context cancel is
// 1. when user makes a request and cancel it by closing the browser or some other way ,we should stop the process the backend is doing
// otherwise unnecessarliy using the resource , etc.,
// 2. when client is making the request and if it takes more time , we should cancel it
// 3. when we have two operations in two go routine and if first go routine fails , we should stop the second go routine
//	otherwise wastage of resources

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/demoListenCancel", listenCancel)
	mux.HandleFunc("/demoEmitCancel", emitCancel)
	mux.HandleFunc("/demoTimeOutCancel", timeoutCancel)
	mux.HandleFunc("/demoDeadlineCancel", deadlineCancel)
	http.ListenAndServe(":8080", mux)
}

//case 1 - Listening for the cancelling event
func listenCancel(writer http.ResponseWriter, request *http.Request) {
	//get the context
	ctx := request.Context()
	// if you dont cancel the request within 10 sec , success will be returned
	//if you close before cancellation is triggered and operations will be suspended
	select {
	case <-time.After(10 * time.Second):
		println("Inside the block after 10 sec")
		writer.Write([]byte("Success"))
	case <-ctx.Done():
		println("the request is cancelled")

	}
}

// case 2 - Emitting the cancel
// when we have two operation if one fails , emit the cancel to other operations
func emitCancel(writer http.ResponseWriter, request *http.Request) {
	//get the context and create the context with cancel
	ctx, cancel := context.WithCancel(request.Context())
	syncWG := sync.WaitGroup{}
	syncWG.Add(2)
	go Operation1(request, &syncWG, cancel)
	go Operation2(request, &syncWG, ctx)
	syncWG.Wait()
	println("End of operations from the mail call")
}

//we need the context with cancel and request context are different
//so both the operations 1 and 2 should have the same context whixh is context with cancel
func Operation2(r *http.Request, group *sync.WaitGroup, ctx context.Context) {
	println("Operation 2 method")
	select {
	case <-time.After(15 * time.Second):
		fmt.Println("Operation performed")
		group.Done()
	case <-ctx.Done(): // here r.context.Done() is different from the ctx becuase this ctx is cancel context
		fmt.Println("canceled operation ")
		group.Done()
	}
}
func Operation1(r *http.Request, group *sync.WaitGroup, cancelContext context.CancelFunc) {
	println("Operation1  method")
	defer group.Done()
	cancelContext() //cancel all the operation using this context
	println("close is called from operation 1 method")
}

//use case 3- time out based context
//if you want to cancel before that time period , you can use like second method cancel
// timeout have cancel future
func timeoutCancel(writer http.ResponseWriter, request *http.Request) {
	ctx, _ := context.WithTimeout(request.Context(), 2*time.Second) // after 2 seconds cancellation will happen
	//if you want you can cancel also
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation performed")
	case <-ctx.Done() :
		fmt.Println("cancelled operations")
	}

}

//use case 3- dead line out based context
//if you want to cancel before that deadline , you can use like second method cancel
// timeout have cancel future
func deadlineCancel(writer http.ResponseWriter, request *http.Request) {
	ctx, _ := context.WithDeadline(request.Context(), time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC)) // after 2 seconds cancellation will happen
	//if you want you can cancel also
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation performed")
	case <-ctx.Done() :
		fmt.Println("cancelled operations")
	}

}
