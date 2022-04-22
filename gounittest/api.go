package main

import (
	"encoding/json"
	"net/http"
)

type UserDB interface {
	FindByID(ID string) (string, error)
}

func init() {

}

func GetUser(db UserDB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		type User struct {
			Name string `json:"name"`
		}
		name, err := db.FindByID("1")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := User{
			Name: name,
		}
		b, _ := json.Marshal(&user)
		rw.Write(b)
		return

	}
}
