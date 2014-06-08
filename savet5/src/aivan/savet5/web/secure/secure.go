package secure

import (
	"aivan/savet5/db/user"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

type SecureHandlerType func(w http.ResponseWriter, r *http.Request, user user.User)
type HandlerType func(w http.ResponseWriter, r *http.Request)

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
			log.Println("session:", session)
			log.Println("user:", us)

			handler(w, r, realUser)
		}
	}
}
