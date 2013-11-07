package main

import(
    "net/http"
    "log"
    "flag"
    "fmt"
)

var hostfolder string
var port string

func main(){
    //The port flag
    portflag := flag.String("port", "3000",
    "The port to host the server on")

    //Host folder
    folder := flag.String("folder", "web",
    "The folder with the webpage to host")

    flag.Parse()

    hostfolder = fmt.Sprintf("%s/", *folder)

    port = fmt.Sprintf(":%s", *portflag)


    r := GetRouter()
    http.Handle("/", r)


    log.Printf("Starting server on port: %s", *portflag)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
