package main

import (
	. "github.com/ivanarandjelovic/savet5/db"
	"github.com/ivanarandjelovic/savet5/web"
	"github.com/ivanarandjelovic/savet5/web/live"
	//"github.com/ivanarandjelovic/savet5/web/secure"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	fmt.Println(DB)
}

var HomeFolder = "static"
var StaticFolders = []string{"html", "js", "css", "images", "fonts"}

func main() {
	fmt.Println("Ziv sam!")
	//DB.DB().Ping()
	fmt.Println(DB)
	if DB.DB().Ping() != nil {
		log.Panicln("No DB connection!")
		//Will exit after this Fatal anyway, no need for return
	}

	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/html/", 301))

	for _, f := range StaticFolders {
		f2 := "/" + f + "/"
		r.PathPrefix(f2).Handler(http.StripPrefix(f2, http.FileServer(http.Dir(HomeFolder+f2))))
	}

	r.HandleFunc("/login", web.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", web.LogoutHandler).Methods("POST")
	r.HandleFunc("/currentUser", web.CurrentUserHandler).Methods("GET")

	r.HandleFunc("/savet", web.SavetHandler).Methods("GET")
	r.HandleFunc("/savet", web.SaveSavetHandler).Methods("POST")
	r.HandleFunc("/savet/{id}", web.GetSavetHandler).Methods("GET")

	r.HandleFunc("/stanari/{savetId}", web.GetStanariHandler).Methods("GET")
	r.HandleFunc("/stanari/{savetId}", web.SaveStanarHandler).Methods("POST")

	r.Handle("/live/getSecured", websocket.Handler(live.WebSocketHandler))

	http.Handle("/", r)

	e := http.ListenAndServe(":8080", context.ClearHandler(r))
	if e != nil {
		println(e.Error())
	}

	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
	fmt.Println("Odoh!")
}
