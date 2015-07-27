package main

import "net/http"

func main() {
	dir := http.FileServer(http.Dir("assets"))
	http.ListenAndServe(":9000", dir)
}

//serve directory binary is in:
/*

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	serveThisDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	dir := http.FileServer(http.Dir(serveThisDir))
	http.ListenAndServe(":9000", dir)
}

*/
