package web

import (
	"aivan/savet5/db/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type loginForm struct {
	Email    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var l loginForm
	err := decoder.Decode(&l)
	if err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintln("LoginHandler called with u:"+l.Email+" and p:(len)", len(l.Password)))

	user, err := user.Login(l.Email, l.Password)

	if err != nil {
		log.Println("returning NotFound for login")
		http.NotFound(w,r)
	} else {
		encoder := json.NewEncoder(w)
		encoder.Encode(user)
		//fmt.Fprint(w, user)
	}
}
