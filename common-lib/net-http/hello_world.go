package main

import (
	"fmt"
	"net/http"
)

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g)
}

func hello(w http.ResponseWriter, r *http.Request)  {
}


type handlerTest struct {}

func (handlerTest *handlerTest) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("ServeHTTP"))
}

func (handlerTest *handlerTest)Test(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Test"))
}

func main() {

	//test.TestTwo()

	http.Handle("/greeting", greeting("Welcome, dj"))

	http.HandleFunc("/Hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hi xiaomi's</h1> "))
	})


	http.Handle("/handlerTest", &handlerTest{})


	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("http server error:", err)
	}


}
