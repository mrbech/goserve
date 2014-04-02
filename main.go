/*
This file is part of Goserve.

    Goserve is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    Goserve is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with Goserve.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

//Location of the host folder in the format: folder/
var hostfolder string

//The port number in the format: :port
var port string

//The index page files, that will be served on "/" requests
var indexfile string

func main() {
	//Setup the commandline flags
	portflag := flag.String("port", "3000",
		"The port to serve to")

	folderflag := flag.String("folder", ".",
		"Folder to serve")

	indexflag := flag.String("index", "index.html",
		"Index file")

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
