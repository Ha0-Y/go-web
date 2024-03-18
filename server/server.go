package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	fmt.Printf("ID: %v\n", id)
	w.Write([]byte(fmt.Sprintf("Hello %v", id)))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Write([]byte(fmt.Sprintf("Register %v", r.Form.Get("name"))))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("upload")
	if err != nil {
		fmt.Fprint(w, "Error uploading")
		return
	}
	data, _ := io.ReadAll(file)
	fmt.Fprintf(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("server/index.html")
	t.Execute(w, "hello")
}

func main() {
	http.HandleFunc("/", http.HandlerFunc(rootHandler))
	http.HandleFunc("/login", http.HandlerFunc(loginHandler))
	http.HandleFunc("/register", http.HandlerFunc(registerHandler))
	http.HandleFunc("/upload", http.HandlerFunc(uploadHandler))
	http.HandleFunc("/hello", http.HandlerFunc(helloHandler))
	http.HandleFunc("/404", http.NotFound)
	http.ListenAndServe(":8080", nil)
}
