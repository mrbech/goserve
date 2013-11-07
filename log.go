package main
import(
    "log"
    "net/http"
)

/*
Log an incomming request
*/
func LogRequest(r *http.Request){
    log.Printf("%s: %s", r.RemoteAddr, r.URL.Path)
}
