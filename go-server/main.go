package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"ParseForm() err : %v",err)
		return
	}

	fmt.Fprintf(w,"POST request successfully")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w,"Name : %s\n",name)
	fmt.Fprintf(w,"Address : %s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/hello" {
		http.Error(w, "path is incorrect", http.StatusBadRequest)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"method not supported",http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w,"Hello World!")
}



func main()  {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err:=http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}