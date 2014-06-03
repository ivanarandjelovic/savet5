package web

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(fmt.Sprintln("LoginHandler called with u:"+l.Email+" and p:(len)", len(l.Password)))

	fmt.Fprint(w, "")
}
