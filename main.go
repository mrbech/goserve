package main

import(
    "net/http"
    "log"
    "flag"
    "fmt"
)

var hostfolder string
var port string
var indexfile string

func main(){
    //The port flag
    portflag := flag.String("port", "3000",
    "The port to serve to, default is 3000")

    //Host folder
    folder := flag.String("folder", ".",
    "Folder to serve, default is the current folder")

    //Index file
    index := flag.String("index", "index.html",
    "Index file, default is index.html")

    flag.Parse()

    hostfolder = fmt.Sprintf("%s/", *folder)

    port = fmt.Sprintf(":%s", *portflag)

    indexfile = *index


    r := GetRouter()
    http.Handle("/", r)


    log.Printf("Starting server on port: %s", *portflag)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
