package ch1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

//Server2 count the number of request,a request to the URL /count return
//the count, other increase count.
func Server2() {
	http.HandleFunc("/", handler2Richer)
	http.HandleFunc("/count", counter)
	//draw lissajous figure
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		Lissajous(w)
	})
	//draw sin figure
	http.HandleFunc("/sin", func(w http.ResponseWriter, r *http.Request) {
		SinGif(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

//handler echoes the path component of the requested URL
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//handler2Richer return more infomation about request.
func handler2Richer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	//go allows a simple statement such as a local variable declaration
	//to precede the if condition, which is particularly useful for error handling.
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
