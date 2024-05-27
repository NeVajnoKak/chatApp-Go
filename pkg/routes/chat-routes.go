package routes

import (
	"github.com/NeVajnoKak/chatApp-Go/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterchatStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/chat/", controllers.CreateChat).Methods("POST")
	router.HandleFunc("/chat/", controllers.GetChat).Methods("GET")
	router.HandleFunc("/chat/{chatId}", controllers.GetChatById).Methods("GET")
	router.HandleFunc("/chat/{chatId}", controllers.UpdateChat).Methods("PUT")
	router.HandleFunc("/chat/{chatId}", controllers.DeleteChat).Methods("DELETE")
}
