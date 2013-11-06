package gobuilder

import(
    "github.com/gorilla/mux"
    "net/http"
)

func FileHandler(rw http.ResponseWriter, r *http.Request){
    LogRequest(r)
    vars := mux.Vars(r)
    file := vars["file"]
    if file == "" {
        http.ServeFile(rw, r, "web/index.html")
    }else{
        http.ServeFile(rw, r, "web/"+file)
    }
}
