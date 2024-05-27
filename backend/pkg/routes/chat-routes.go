package routes

import (
	"github.com/NeVajnoKak/chatApp-Go/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterChatRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/chat/", controllers.CreateChat).Methods("POST")
	router.HandleFunc("/api/chat/", controllers.GetChat).Methods("GET")
	router.HandleFunc("/api/chat/{chatId}", controllers.GetChatById).Methods("GET")
	router.HandleFunc("/api/chat/{chatId}", controllers.UpdateChat).Methods("PUT")
	router.HandleFunc("/api/chat/{chatId}", controllers.DeleteChat).Methods("DELETE")
}
