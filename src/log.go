package gobuilder
import(
    "log"
    "net/http"
)

func LogRequest(r *http.Request){
    log.Printf("%s: %s", r.RemoteAddr, r.URL.Path)
}
