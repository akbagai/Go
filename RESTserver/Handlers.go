package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("mssql", "server=myserver;user id=domain\\username;password=itssecret;database=myDB")
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
	rows, err := db.Query("select * from TABLE")
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

	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		panic(err)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func IndexPOST(w http.ResponseWriter, r *http.Request) {
	/*
		r.ParseForm()
		fmt.Fprintln(w, r.Form.Get("login")+" "+r.Form.Get("password"))
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("mssql", "server=myserver;user id=domain\\username;password=itssecret;database=myDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r.ParseForm()

	rows, err := db.Prepare("insert into TABLE values (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := rows.Exec(r.Form.Get("id"), r.Form.Get("name"), r.Form.Get("age"), r.Form.Get("gender"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "200 OK Buddy")
	affected, _ := res.RowsAffected()
	fmt.Fprintln(w, affected)
}

func IndexDELETE(w http.ResponseWriter, r *http.Request) {
	/*
		
		variables := mux.Vars(r)
		v := variables["id"]
		fmt.Fprintln(w, v)
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("mssql", "server=myserver;user id=domain\\username;password=itssecret;database=myDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Prepare("delete from TABLE where id=?")
	if err != nil {
		log.Fatal(err)
	}
	variables := mux.Vars(r)
	v := variables["id"]
	res, err := rows.Exec(v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "200 OK Buddy")
	affected, _ := res.RowsAffected()
	fmt.Fprintln(w, affected)
}

func IndexPUT(w http.ResponseWriter, r *http.Request) {

	/*
	variables := mux.Vars(r)
	v := variables["id"]
	r.ParseForm()
	fmt.Fprintln(w, v)
	fmt.Fprintln(w, r.Form.Get("login"))
	*/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("mssql", "server=myserver;user id=domain\\username;password=itssecret;database=myDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r.ParseForm()

	rows, err := db.Prepare("update TABLE set name=?, age=?, gender=? where id=?")
	if err != nil {
		log.Fatal(err)
	}
	variables := mux.Vars(r)
	v := variables["id"]
	r.ParseForm()
	res, err := rows.Exec(r.Form.Get("name"), r.Form.Get("age"), r.Form.Get("gender"), v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "200 OK Buddy")
	affected, _ := res.RowsAffected()
	fmt.Fprintln(w, affected)

}
