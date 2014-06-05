package main

import (
	. "aivan/savet5/db"
	"aivan/savet5/web"
	"fmt"
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
	if(DB.DB().Ping() != nil) {
		log.Panicln("No DB connection!")
		//Will exit after this Fatal anyway, no need for return
	}

	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/html/", 301))

	for _, f := range StaticFolders {
		f2 := "/"+f+"/"
		r.PathPrefix(f2).Handler(http.StripPrefix(f2, http.FileServer(http.Dir(HomeFolder+f2))))
	}

	r.HandleFunc("/login", web.LoginHandler).Methods("POST")

	http.Handle("/", r)

	e := http.ListenAndServe(":8080", r)
	if e != nil {
		println(e.Error())
	}

	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
	fmt.Println("Odoh!")
}
