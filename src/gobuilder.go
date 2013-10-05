package main

import(
    "github.com/gorilla/mux"
    "net/http"
    "log"
)
func LogRequest(r *http.Request){
    log.Printf("%s: %s", r.RemoteAddr, r.URL.Path)
}

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

func main(){
    r := mux.NewRouter()
    r.HandleFunc("/{file:.*}", FileHandler).Methods("GET")
    http.Handle("/", r)
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
