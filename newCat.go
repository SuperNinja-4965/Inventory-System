package main

import (
	"net/http"
)

func initNew() {
	http.HandleFunc("/newCategory/", newCategory)
	http.HandleFunc("/newItem/", newItem)
}

func newCategory(w http.ResponseWriter, r *http.Request) {

}

func newItem(w http.ResponseWriter, r *http.Request) {

}
