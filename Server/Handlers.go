package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
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
	content, err := ioutil.ReadFile(os.Getenv("JSONServerFile"))
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
	content, err := ioutil.ReadFile(os.Getenv("JSONServerFile"))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(content)
	var data query
	json.Unmarshal(content, &data)
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	var output []features

	for i := 0; i < len(data.Features); i++ {
		if data.Features[i].Properties[key] == value {
			output = append(output, data.Features[i])

		}

	}
	fmt.Fprint(w, json.NewEncoder(w).Encode(output))

}

func Subset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	content, err := ioutil.ReadFile(os.Getenv("JSONServerFile"))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(content)
	var data query
	json.Unmarshal(content, &data)
	vars := mux.Vars(r)
	howMany, _ := strconv.Atoi(vars["howMany"])
	var output []features
	for i := 0; i < howMany; i++ {
		output = append(output, data.Features[i])

	}
	fmt.Fprint(w, json.NewEncoder(w).Encode(output))
}

func MSsql(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("mssql", "server=myServer;user id=domain\\myName;password=secret;database=myDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id     int
		name   string
		age    string
		gender string
	)

	type Result struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Age    string `json:"age"`
		Gender string `json:"gender"`
	}
	//rows, err := db.Query("select id, name, age, gender from PaulTest where id = ?", 3)
	rows, err := db.Query("select * from PaulTest")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var output []Result
	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &gender)
		if err != nil {
			log.Fatal(err)
		}

		d := Result{id, name, age, gender}
		output = append(output, d)
		fmt.Fprint(w, json.NewEncoder(w).Encode(d))

	}
	fmt.Fprint(w, json.NewEncoder(w).Encode(output))
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
