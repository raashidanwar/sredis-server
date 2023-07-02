package main

import (
	"fmt"
	"net/http"
)

var store = make(map[string]string)

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}
	key := r.URL.Query().Get("key")
	value := store[key]
	if value == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(key + " => " + value))
	return
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	store[key] = value
	w.WriteHeader(201)
	w.Write([]byte(key + " => " + value))
	return
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/set", setHandler)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
