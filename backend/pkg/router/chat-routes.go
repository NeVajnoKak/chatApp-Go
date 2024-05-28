package router

import (
	"backend/pkg/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/chat/{id}", controllers.Getchat).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/chat", controllers.GetAllchat).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newchat", controllers.Createchat).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/chat/{id}", controllers.Updatechat).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletechat/{id}", controllers.Deletechat).Methods("DELETE", "OPTIONS")

	return router
}
