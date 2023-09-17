package main

import (
	"log"
	"net/http"
)

func NoCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/styles/", NoCache(http.StripPrefix("/styles/", http.FileServer(http.Dir("public")))))

	http.Handle("/sidebar", NoCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/components/layout/sidebar.html")
	})))
	
	http.Handle("/header", NoCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/components/layout/header.html")
	})))
	
	http.Handle("/dashboard", NoCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/components/layout/dashboard.html")
	})))
	
	http.Handle("/", NoCache(fs))

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}