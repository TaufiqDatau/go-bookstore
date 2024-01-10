package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "gorm.io/gorm"
	"github.com/TaufiqDatau/go-bookstore/pkg/routes"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.fatal(http.ListenAndServe("localhost:9010",r))
)
