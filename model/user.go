package controllers

import (
	"log"
	"time"

	"github.com/go-pg/pg/v9"
	guuid "github.com/google/uuid"
)

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

type User struct {
	ID        guuid.UUID `json:"id"`
	username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func GetAllUsers() (users []User) {
	err := dbConnect.Model(&users).Select()

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
	}
	return
}

func CreateUser() (user User) {
	email := user.Email
	insertError := dbConnect.Insert(&User{
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if insertError != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", insertError)
	}
	return
}

func GetSingleUser(userId guuid.UUID) (user User) {
	user = User{ID: userId}
	err := dbConnect.Select(user)

	if err != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", err)
	}
	return
}

func EditUser(userId string, m map[string]string) (user User) {
	_, err := dbConnect.Model(&User{}).Set("email = ?", m["email"]).Where("id = ?", userId).Update()

	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
	}
	return
}

func DeleteUser(userId guuid.UUID) (user User) {
	user = User{ID: userId}
	err := dbConnect.Delete(user)

	if err != nil {
		log.Printf("Error while deleting a single user, Reason: %v\n", err)
	}
	return
}
