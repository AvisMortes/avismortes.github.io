package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//http.HandleFunc("/transactions/129562341-baf56f8476", handler3)
	//http.HandleFunc("/ctp/api/checkouts/e3596a225bf79b111ec885a0f07e5f465416b91248c89f892e9f984cdc7681a1", handler2)
	http.HandleFunc("/", handler1)
	http.HandleFunc("/about", aboutHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler1(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method, r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))
	file, err := os.Open("responce.txt")
	defer file.Close()

	data := make([]byte, 64)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	time.Sleep(10 * time.Second)
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break
		}
		fmt.Fprintf(w, "%s", string(data[:n]))
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method, r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))
	file, err := os.Open("responce2.txt")
	defer file.Close()

	data := make([]byte, 64)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break
		}
		fmt.Fprintf(w, "%s", string(data[:n]))
	}
}

func handler3(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method, r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))
	file, err := os.Open("responce3.txt")
	defer file.Close()

	data := make([]byte, 64)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break
		}
		fmt.Fprintf(w, "%s", string(data[:n]))
	}
}
