package gobuilder

import(
    "github.com/gorilla/mux"
    "bitbucket.org/gobuilder/server/filehandler"
)

func GetRouter(hostfolder string) *mux.Router{
    r := mux.NewRouter()
    r.HandleFunc(hostfolder, "/{file:.*}", FileHandler).Methods("GET")

    return r
}
