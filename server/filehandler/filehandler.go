package gobuilder

import(
    "github.com/gorilla/mux"
    "net/http"
    "bitbucket.org/gobuilder/server/log"
)

func FileHandler(hostfolder string, rw http.ResponseWriter, r *http.Request){
    LogRequest(r)
    vars := mux.Vars(r)
    file := vars["file"]
    if file == "" {
        http.ServeFile(rw, r, hostfolder+"index.html")
    }else{
        http.ServeFile(rw, r, hostfolder+file)
    }
}
