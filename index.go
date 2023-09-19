package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/api/styles/", http.StripPrefix("/api/styles/", http.FileServer(http.Dir("public"))))

	http.Handle("/api/sidebar", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/components/layout/sidebar.html")
	}))

	http.Handle("/api/header", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/components/layout/header.html")
	}))

	http.Handle("/api/dashboard", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/components/layout/dashboard.html")
	}))

	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
