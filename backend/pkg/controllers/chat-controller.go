package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected to postgres")
	return db
}

func Createchat(w http.ResponseWriter, r *http.Request) {
	var chat models.chat

	err := json.NewDecoder(r.Body).Decode(&chat)

	if err != nil {
		log.Fatal("Unabel to decide the request body. %v", err)
	}

	insertID := insertchat(chat)

	res := response{
		ID:      insertID,
		Message: "chat created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func Getchat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	chat, err := getchat(int64(id))

	if err != nil {
		log.Fatalf("Unable to get chat. %v", err)
	}

	json.NewEncoder(w).Encode(chat)
}

func GetAllchat(w http.ResponseWriter, r *http.Request) {
	chats, err := getAllchats()

	if err != nil {
		log.Fatalf("Unable to get all the chats. %v", err)
	}

	json.NewEncoder(w).Encode(chats)
}

func Updatechat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	var chat models.chat

	err = json.NewDecoder(r.Body).Decode(&chat)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updateRows := updatechat(int64(id), chat)

	msg := fmt.Sprintf("chats update successfully. Total rows /record affected %v", updateRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func Deletechat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	deletedRows := deletechat(int64(id))

	msg := fmt.Sprintf("chat deleted successfully. Total rows /records %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func insertchat(chat models.chat) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO chats(name, price, company) VALUES ($1 , $2 , $3) RETURNING chatid`
	var id int64

	err := db.QueryRow(sqlStatement, chat.Name, chat.Price, chat.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func getchat(id int64) (models.chat, error) {
	db := CreateConnection()

	defer db.Close()

	var chat models.chat

	sqlStatement := `SELECT * FROM chats WHERE chatid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&chat.chatID, &chat.Name, &chat.Price, &chat.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned!")
		return chat, nil
	case nil:
		return chat, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return chat, err
}

func getAllchats() ([]models.chat, error) {
	db := CreateConnection()

	defer db.Close()

	var chats []models.chat

	sqlStatement := `SELECT * FROM chats`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var chat models.chat
		err = rows.Scan(&chat.chatID, &chat.Name, &chat.Price, &chat.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		chats = append(chats, chat)
	}
	return chats, err
}

func updatechat(id int64, chat models.chat) int64 {
	db := CreateConnection()

	defer db.Close()

	sqlStatement := `UPDATE chats SET name=$2 , price=$3, company=$4 WHERE chatid =$1`
	res, err := db.Exec(sqlStatement, id, chat.Name, chat.Price, chat.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/records affected %v", rowsAffected)

	return rowsAffected
}

func deletechat(id int64) int64 {
	db := CreateConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM chats WHERE chatid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/records affected %v", rowsAffected)

	return rowsAffected
}
