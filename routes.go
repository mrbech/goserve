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

import(
    "github.com/gorilla/mux"
)

/*
Setup the Router
*/
func GetRouter() *mux.Router{
    r := mux.NewRouter()
    r.HandleFunc("/{file:.*}", FileHandler).Methods("GET")

    return r
}
