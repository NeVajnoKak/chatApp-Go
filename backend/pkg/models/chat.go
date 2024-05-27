package models

import (
	"log"
	"time"

	"github.com/NeVajnoKak/chatApp-Go/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Chat struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Message   string     `json:"message"`
	UserId    string     `json:"userId"`
	// Publication string     `json:"publication"`
}

func init() {
	var err error
	config.Connect()
	db = config.GetDb()
	if db == nil {
		log.Fatalf("Failed to connect to the database.")
	}
	if err = db.AutoMigrate(&Chat{}).Error; err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func (b *Chat) CreateChat() (*Chat, error) {
	if err := db.Create(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllChats() ([]Chat, error) {
	var Chats []Chat
	if err := db.Find(&Chats).Error; err != nil {
		return nil, err
	}
	return Chats, nil
}

func GetChatById(Id int64) (*Chat, *gorm.DB, error) {
	var Chat Chat
	db := db.Where("id = ?", Id).Find(&Chat)
	if db.Error != nil {
		return nil, nil, db.Error
	}
	return &Chat, db, nil
}

func DeleteChat(ID int64) error {
	if err := db.Where("id = ?", ID).Delete(&Chat{}).Error; err != nil {
		return err
	}
	return nil
}
