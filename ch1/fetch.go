package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Fetch fetch url from command arguments, and print reply.
func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}

//FetchExe1_7 fetch url from command arguments, and print reply using io.Copy(dst, src)
//and os.Stdout
func FetchExe1_7() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		byteNumber, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("recevie byte number : %d\n", byteNumber)
	}
}

//FetchExe1_8 fetch url from command arguments, and print reply using io.Copy(dst, src)
//and os.Stdout, check if have http:// , if not ,add it.
func FetchExe1_8() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		byteNumbers, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("recevie byte numbers : %d\n", byteNumbers)
	}
}

//FetchExe1_9 fetch url from command arguments, and print reply using io.Copy(dst, src)
//and os.Stdout, check if have http:// , if not ,add it. add print status code.
func FetchExe1_9() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		byteNumbers, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: status code %d reading %s : %v\n", resp.StatusCode, url, err)
			os.Exit(1)
		}
		fmt.Printf("recevie byte numbers : %d, status code : %d\n", byteNumbers, resp.StatusCode)
	}
}
