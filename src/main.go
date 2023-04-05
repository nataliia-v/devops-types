package main

import "net/http"

func main() {
	http.ListenAndServe(":8082", http.FileServer(http.Dir("./html")))
} 
