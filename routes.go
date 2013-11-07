package main

import(
    "github.com/gorilla/mux"
)

/*
Setup the Router
*/
func GetRouter() *mux.Router{
    r := mux.NewRouter()
    r.HandleFunc("/{file:.*}", FileHandler).Methods("GET")

    return r
}
