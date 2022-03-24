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
	http.HandleFunc("/", handler)
	http.HandleFunc("/rss/category/now.json", getInfo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func getData(name string) []byte {
	file, err := os.Open(name)
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
		return data[:n]
	}
	return data
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method, r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))
	data := getData("response.txt")
	fmt.Fprintf(w, "%s", string(data))
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.UserAgent(), r.RemoteAddr, time.Now().Format(time.RFC3339))
	data := getData("getInfo.txt")
	fmt.Fprintf(w, "%s", string(data))
}
