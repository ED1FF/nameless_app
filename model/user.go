package model

import (
	"fmt"
	"log"
	"net/url"
	"time"

	crypto "nameless_app/lib"

	"github.com/go-pg/pg/v9"
	guuid "github.com/google/uuid"
)

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

type User struct {
	ID        guuid.UUID `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (user *User) buildUser(params url.Values) *User {
	crypto_password := crypto.Generate(params.Get("password"))

	user.Email = params.Get("email")
	user.Username = params.Get("username")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if crypto_password != "" {
		user.Password = crypto_password
	}

	return user
}

func CreateUser(params url.Values) (user User, err string) {
	insertError := dbConnect.Insert(user.buildUser(params))

	if insertError != nil {
		err = fmt.Sprintf("Error while inserting new user into db, Reason: %v\n", insertError)
	}
	return
}

func GetAllUsers() (users []User) {
	err := dbConnect.Model(&users).Select()

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
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
