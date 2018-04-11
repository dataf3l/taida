package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	//"strings"
)

func update_repo(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()  // parse arguments, you have to call this by yourself
	//fmt.Println(r.Form)  // print form information in server side
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//    fmt.Println("key:", k)
	//    fmt.Println("val:", strings.Join(v, ""))
	//}

	out, err := exec.Command("/usr/bin/git", "pull", "origin", "master").Output()
	if err != nil {
		fmt.Fprintf(w, string(err.Error())) // send data to client side
		return

	}

	fmt.Fprintf(w, string(out)) // send data to client side
}

func main() {
	http.HandleFunc("/update_repository", update_repo) // set router
	err := http.ListenAndServe(":6520", nil)           // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
