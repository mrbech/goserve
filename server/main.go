package main

import(
    "net/http"
    "log"
    "flag"
    "fmt"
    "bitbucket.org/gobuilder/server/routes"
    "bitbucket.org/gobuilder/server/log"
)

func main(){
    //The port flag
    portflag := flag.String("port", "3000",
    "The port to host the server on, default: 3000")

    //Host folder
    folder := flag.String("folder", "web",
    "The folder to host the webpage from default: web")

    hostfolder := fmt.Sprintf("%s/",folder)

    port := fmt.Sprintf(":%s", *portflag)


    r := GetRouter(hostfolder)
    http.Handle("/", r)

    flag.Parse()

    log.Printf("Starting server on port: %s", *portflag)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
