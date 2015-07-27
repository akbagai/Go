package main

import "net/http"

func main() {
	dir := http.FileServer(http.Dir("assets"))
	http.ListenAndServe(":9000", dir)
}
