package gobuilder

import(
    "github.com/gorilla/mux"
)

func GetRouter() *mux.Router{
    r := mux.NewRouter()
    r.HandleFunc("/{file:.*}", FileHandler).Methods("GET")

    return r
}

