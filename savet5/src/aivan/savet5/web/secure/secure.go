package secure

import (
	"aivan/savet5/db/user"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"code.google.com/p/go.net/websocket"
)

type SecureHandlerType func(w http.ResponseWriter, r *http.Request, user user.User)
type HandlerType func(w http.ResponseWriter, r *http.Request)

type SecureWSHandlerType func (ws *websocket.Conn, user user.User)
type WSHandlerType func (ws *websocket.Conn)

/*
	Wrap regular handler in a check if user has valid session, and provide the user to the wrapped handler
*/
func SecureHandler(store *sessions.FilesystemStore, handler SecureHandlerType) HandlerType {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("SecureHandler checking for user")

		session, _ := store.Get(r, "XSRF-TOKEN")

		log.Println("session:", session)

		var us = session.Values["user"]

		if us == nil {
			log.Println("No user in session")
			http.Error(w, "401 No current user", 401)
		} else {
			realUser := us.(user.User)
			handler(w, r, realUser)
		}
	}
}

/* 
	Wrap regular handler in a check if user has valid session, and provide the user to the wrapped handler
*/
func SecureWSHandler(store *sessions.FilesystemStore, handler SecureWSHandlerType) WSHandlerType {
	return func (ws *websocket.Conn) {

		log.Println("SecureWSHandler checking for user")

		session, _ := store.Get(ws.Request(), "XSRF-TOKEN")

		log.Println("session:", session)

		var us = session.Values["user"]

		if us == nil {
			log.Println("No user in session. Closing websocket!")
			ws.Close()
		} else {
			realUser := us.(user.User)
			handler(ws, realUser)
		}
	}
}
