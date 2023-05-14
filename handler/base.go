package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tirthankarkundu17/ecommerce-price-checker/db"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {

	db, err := db.NewDB()
	if err != nil {
		log.Fatal("Error while connecting to DB")
	}
	err = db.AutoMigrate()
	if err != nil {
		log.Fatal("Error while migrating to DB", err)
	}
	server.DB = db.Conn

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr, port string) {
	fmt.Println("Listening to port " + port)
	log.Fatal(http.ListenAndServe(addr+":"+port, server.Router))
}
