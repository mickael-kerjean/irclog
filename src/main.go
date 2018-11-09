package main

import (
	"github.com/mickael-kerjean/mux"
	"github.com/mickael-kerjean/irc/src/ctrl"
	"html/template"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/search").Methods("GET")
	api.HandleFunc("/messages").Methods("GET")

	r.HandleFunc("/assets").Handler
	r.HandleFunc("/", IndexHandler).Methods("GET")	

	src := &http.Server{
		Addr:    ":8335",
		Handler: r,
	}
	//Log.Info("STARTING SERVER")
	if err := srv.ListenAndServe(); err != nil {
		return
	}
}
