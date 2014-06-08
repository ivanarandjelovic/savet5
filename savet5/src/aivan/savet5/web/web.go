package web

import (
	"aivan/savet5/db/savet"
	"aivan/savet5/db/stanari"
	"aivan/savet5/db/user"
	"aivan/savet5/web/secure"
	"encoding/json"
	"fmt"
	//"github.com/dchest/uniuri"
	"encoding/gob"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"strconv"
)

var Store = sessions.NewFilesystemStore("", securecookie.GenerateRandomKey(32))

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
		log.Println("returning NotFound for login")
		http.NotFound(w, r)
	} else {
		// User logged in OK:
		// Get a session.
		//sessKey := uniuri.New()
		session, _ := Store.Get(r, "XSRF-TOKEN")
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

var CurrentUserHandler = secure.SecureHandler(Store, func(w http.ResponseWriter, r *http.Request, user user.User) {
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
})

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Logout called")

	session, _ := Store.Get(r, "XSRF-TOKEN")

	log.Println("session:", session)

	//Delete session
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		log.Println("Session save failed during logout! message: ", err)
	}
}

var SavetHandler = secure.SecureHandler(Store, func(w http.ResponseWriter, r *http.Request, user user.User) {
	log.Println("SavetHandlerImpl called")
	saveti, _ := savet.List()
	encoder := json.NewEncoder(w)
	encoder.Encode(saveti)
})

//var SavetHandler = secure.SecureHandler(Store, w, r , SavetHandlerImpl)

var SaveSavetHandler = secure.SecureHandler(Store, func(w http.ResponseWriter, r *http.Request, user user.User) {
	decoder := json.NewDecoder(r.Body)
	var s savet.Savet
	err := decoder.Decode(&s)
	if err != nil {
		panic(err)
	}

	log.Println("SaveSavetHandler called with savet:", s)

	err = savet.Create(s)

	if err != nil {
		http.NotFound(w, r)
	}
})

var GetSavetHandler = secure.SecureHandler(Store, func(w http.ResponseWriter, r *http.Request, user user.User) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	savet := savet.Get(id)
	encoder := json.NewEncoder(w)
	encoder.Encode(savet)
})

var GetStanariHandler = secure.SecureHandler(Store, func(w http.ResponseWriter, r *http.Request, user user.User) {
	vars := mux.Vars(r)
	savetId, _ := strconv.ParseInt(vars["savetId"], 10, 64)
	savet, err := stanari.Get(savetId)
	if err != nil {
		log.Println("Error", err)
		http.NotFound(w, r)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(savet)
})
