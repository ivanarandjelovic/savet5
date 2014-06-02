package main

import (
	. "aivan/savet5/db"
	"aivan/savet5/web"
	"fmt"
	"github.com/gorilla/mux"
	//	"log"
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

	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/html/", 301))

	for _, f := range StaticFolders {
		f2 := "/"+f+"/"
		r.PathPrefix(f2).Handler(http.StripPrefix(f2, http.FileServer(http.Dir(HomeFolder+f2))))
	}

	//r.PathPrefix("/html/").Handler(http.StripPrefix("/html/", http.FileServer(http.Dir(HomeFolder+"html/"))))
	//r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(HomeFolder+"js/"))))
	//r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(HomeFolder+"css/"))))

	//r.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("static/html"))))
	//r.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	//r.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	//r.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("wwwroot"))))
	//r.HandleFunc("/web", web.HomeHandler).Methods("GET", "POST")

	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)

	r.HandleFunc("/login", web.LoginHandler).Methods("POST")


	http.Handle("/", r)

	e := http.ListenAndServe(":8080", r)
	if e != nil {
		println(e.Error())
	}

	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
	fmt.Println("Odoh!")
}
