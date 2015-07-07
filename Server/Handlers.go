package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type query struct {
	Features []features
}

type features struct {
	Properties map[string]interface{}
}

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Welcome!")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	content, err := ioutil.ReadFile("C:/Users/Me/Desktop/GoCode/src/Server/files/art.json")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(content)
	var data query
	json.Unmarshal(content, &data)
	if err := json.NewEncoder(w).Encode(data.Features); err != nil {
		panic(err)
	}
}
