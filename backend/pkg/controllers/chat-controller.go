package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NeVajnoKak/chatApp-Go/pkg/models"
	"github.com/NeVajnoKak/chatApp-Go/pkg/utils"
	"github.com/gorilla/mux"
)

func GetChat(w http.ResponseWriter, r *http.Request) {
	newChats, err := models.GetAllChats()
	if err != nil {
		http.Error(w, "Error fetching Chats: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(newChats)
	if err != nil {
		http.Error(w, "Error marshalling Chats: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetChatById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ChatId := vars["ChatId"]
	ID, err := strconv.ParseInt(ChatId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Chat ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	ChatDetails, _, err := models.GetChatById(ID)
	if err != nil {
		http.Error(w, "Error fetching Chat: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(ChatDetails)
	if err != nil {
		http.Error(w, "Error marshalling Chat: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateChat(w http.ResponseWriter, r *http.Request) {
	CreateChat := &models.Chat{}
	utils.ParseBody(r, CreateChat)
	b, err := CreateChat.CreateChat()
	if err != nil {
		http.Error(w, "Error creating Chat: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Error marshalling Chat: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ChatId := vars["ChatId"]
	ID, err := strconv.ParseInt(ChatId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Chat ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = models.DeleteChat(ID)
	if err != nil {
		http.Error(w, "Error deleting Chat: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateChat(w http.ResponseWriter, r *http.Request) {
	updateChat := &models.Chat{}
	utils.ParseBody(r, updateChat)
	vars := mux.Vars(r)
	ChatId := vars["ChatId"]
	ID, err := strconv.ParseInt(ChatId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Chat ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	ChatDetails, db, err := models.GetChatById(ID)
	if err != nil {
		http.Error(w, "Error fetching Chat: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if updateChat.Message != "" {
		ChatDetails.Message = updateChat.Message
	}
	if updateChat.UserId != "" {
		ChatDetails.UserId = updateChat.UserId
	}
	if err := db.Save(&ChatDetails).Error; err != nil {
		http.Error(w, "Error updating Chat: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(ChatDetails)
	if err != nil {
		http.Error(w, "Error marshalling Chat: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
