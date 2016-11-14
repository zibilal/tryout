package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"html/template"
)

type add struct {
	Sum int
}

var templates = template.Must(template.ParseFiles("/Users/zibilal/Documents/learning-curves/go/src/github.com/zibilal/tryout/gotestbench/template/template.html"))

func handleStructAdd(w http.ResponseWriter, r *http.Request) {

	first, second := r.FormValue("first"), r.FormValue("second")
	one, err := strconv.Atoi(first)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	two, err := strconv.Atoi(second)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	m := struct{a, b int} {one, two}
	structSum := add{Sum: m.a + m.b}


	err = templates.Execute(w, structSum)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func main() {
	fmt.Println("Starts...")

	http.HandleFunc("/struct", handleStructAdd)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
