package main

import(
    "github.com/gorilla/mux"
    "net/http"
)

func FileHandler(rw http.ResponseWriter, r *http.Request){
    LogRequest(r)
    vars := mux.Vars(r)
    file := vars["file"]
    if file == "" {
        http.ServeFile(rw, r, hostfolder+index)
    }else{
        http.ServeFile(rw, r, hostfolder+file)
    }
}
