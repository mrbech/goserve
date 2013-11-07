package main

import(
    "net/http"
    "log"
    "flag"
    "fmt"
)

var hostfolder string
var port string
var index string

func main(){
    //The port flag
    portflag := flag.String("port", "3000",
    "The port to host the server on, default is 3000")

    //Host folder
    folder := flag.String("folder", ".",
    "Folder to host, default is current folder")

    //Index file
    index = *flag.String("index", "index.html",
    "Index file, default is index.html")

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
