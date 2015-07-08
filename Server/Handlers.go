package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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
	content, err := ioutil.ReadFile("C:/Users/me/Desktop/GoCode/src/server/files/art.json")
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

func FilterQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	content, err := ioutil.ReadFile("C:/Users/me/Desktop/GoCode/src/server/files/art.json")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(content)
	var data query
	json.Unmarshal(content, &data)
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	for i := 0; i < len(data.Features); i++ {
		if data.Features[i].Properties[key] == value {
			if err := json.NewEncoder(w).Encode(data.Features[i]); err != nil {
				panic(err)
			}
		} else {
		}
	}

}
