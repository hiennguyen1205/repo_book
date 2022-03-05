package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	routes "github.com/hiennguyen1205/repo_book/pkg/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
