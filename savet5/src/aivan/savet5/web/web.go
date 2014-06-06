package web

import (
	"aivan/savet5/db/user"
	"encoding/json"
	"fmt"
	//"github.com/dchest/uniuri"
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store = sessions.NewFilesystemStore("", securecookie.GenerateRandomKey(32))

func init() {
	gob.Register(user.User{})
}

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
		//sessKey := uniuri.New()
		session, _ := store.Get(r, "XSRF-TOKEN")
		session.Values["user"] = user
		log.Println("Session value user set to:", session.Values["user"])
		//cookie := sessions.NewCookie("XSRF-TOKEN", sessKey, &(sessions.Options{Path: "/", MaxAge: 30 * 60, HttpOnly: true}))
		//http.SetCookie(w, cookie)
		err := session.Save(r, w)
		if err != nil {
			log.Println("Session save failed! message: ", err)
		}
		log.Println("session:", session)
		encoder := json.NewEncoder(w)
		encoder.Encode(tokenStruct{session.ID})
		//fmt.Fprint(w, user)
	}
}

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Current user called")

	session, _ := store.Get(r, "XSRF-TOKEN")

	log.Println("session:", session)

	user := session.Values["user"]

	if user == nil {
		log.Println("No user in session")
		http.Error(w, "401 No current user", 401)
	} else {

		log.Println("session:", session)
		log.Println("user:", user)

		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Logout called")

	session, _ := store.Get(r, "XSRF-TOKEN")

	log.Println("session:", session)

	//Delete session
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		log.Println("Session save failed during logout! message: ", err)
	}

}
