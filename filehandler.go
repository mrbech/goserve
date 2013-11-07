package main

import(
    "github.com/gorilla/mux"
    "net/http"
)

/*
Handles the file serving
*/
func FileHandler(rw http.ResponseWriter, r *http.Request){
    LogRequest(r)
    vars := mux.Vars(r)
    file := vars["file"]
    if file == "" {
        http.ServeFile(rw, r, hostfolder+indexfile)
    }else{
        http.ServeFile(rw, r, hostfolder+file)
    }
}
