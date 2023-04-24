package main

import (
	"log"
	"net/http"

	"github.com/Yang-Shun-Yu/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()                                //create a new instance of the mux router,which will be used to handle incoming HTTP requests
	routes.RegisterBookStoreRoutes(r)                   //call a function RegisterBookStoreRoutes which registers all the routes(URL endpoints) for the book store application with the router
	http.Handle("/", r)                                 //register the mux router as the handler for all incoming HTTP requests to the root URL(/)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) //start the http server on port 9010 of localhost interface. if an error occurs while starting the server,it logs the error and exits the program

}
