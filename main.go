package main

import(
    "net/http"
    "log"
    "flag"
    "fmt"
)

//Gobal config variables
var hostfolder string
var port string
var indexfile string

func main(){
    //Setup the commandline flags
    portflag := flag.String("port", "3000",
    "The port to serve to, default is 3000")

    folderflag := flag.String("folder", ".",
    "Folder to serve, default is the current folder")

    indexflag := flag.String("index", "index.html",
    "Index file, default is index.html")

    flag.Parse()

    //Setup the shared config variables
    hostfolder = fmt.Sprintf("%s/", *folderflag)
    port = fmt.Sprintf(":%s", *portflag)
    indexfile = *indexflag


    //Setup the routing
    r := GetRouter()
    http.Handle("/", r)


    //Start the server
    log.Printf("Starting server on port: %s", *portflag)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("Could not start the server: ", err)
    }
}
