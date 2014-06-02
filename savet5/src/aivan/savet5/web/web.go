package web

import (
	"fmt"
	"net/http"
	"log"
)

type loginForm struct {
    username string
    password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	log.Println(r.PostForm)
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println(r.PostForm)

	fmt.Println(fmt.Sprintln("LoginHandler called with u:"+username+" and p:(len)", len(password)))

	fmt.Fprint(w, "")
}
