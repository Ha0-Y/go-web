package main

import (
	`fmt`
	`io`
	`log`
	`net/http`
)

func testGet() {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("Error get url: %v", err)
		return
	}
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	fmt.Printf("%s", string(b))
}

func testHeader() {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("Error get url: %v", err)
		return
	}
	defer r.Body.Close()
	fmt.Println(r.Header)
	fmt.Println(r.Header.Get("User-Agent"))
}

func main() {
	testHeader()
}
