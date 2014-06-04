package web

import (
	"aivan/savet5/db/user"
	"encoding/json"
	"fmt"
	"github.com/dchest/uniuri"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("molto-secreto"))

type loginForm struct {
	Email    string
	Password string
}

type tokenStruct struct {
	Token string `json:"token"`
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
		http.NotFound(w, r)
	} else {
		// User logged in OK:
		// Get a session.
		sessKey := uniuri.New()
		session, _ := store.Get(r, sessKey)
		session.Values["user"] = user
		cookie := sessions.NewCookie("XSRF-TOKEN", sessKey, &(sessions.Options{Path: "/", MaxAge: 30 * 60, HttpOnly: true}))
		http.SetCookie(w, cookie)
		session.Save(r, w)
		encoder := json.NewEncoder(w)
		encoder.Encode(tokenStruct{sessKey})
		//fmt.Fprint(w, user)
	}
}

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	
	encoder := json.NewEncoder(w)
	

	log.Println("Current user called")

	token := r.FormValue("token")

	if len(token) == 0 {
		log.Println("No token!")
		encoder.Encode(user.User{})
	} else {

		log.Println("token:", token)

		session, _ := store.Get(r, token)
		
		log.Println("user:", session.Values["user"])

		encoder.Encode(session.Values["user"])
	}
}
