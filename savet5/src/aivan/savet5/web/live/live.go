package live

import (
	"aivan/savet5/db/user"
	"aivan/savet5/web"
	"aivan/savet5/web/secure"
	"code.google.com/p/go.net/websocket"
	//	"io"
	"fmt"
	"log"
)

var counter func() int

func init() {
	counter = func() func() int {
		counter := 0
		return func() int {
			log.Println("c je ", counter)
			counter++
			return counter
		}
	}()
}

var WebSocketHandler = secure.SecureWSHandler(web.Store, func(ws *websocket.Conn, user user.User) {
	log.Println("New websocket connection!")

	for {
		var text string
		err := websocket.Message.Receive(ws, &text)
		if err != nil {
			fmt.Println(err)
			break
		}
		log.Println("Received WS message: ", text)
		currVal := counter()
		log.Println("counter call:", currVal)
		text = fmt.Sprint("Response: ", currVal)
		log.Println("Sending response: ", text)
		err = websocket.Message.Send(ws, text)
		if err != nil {
			fmt.Println(err)
			break
		}
		if currVal%5 == 0 {
			ws.Close()
			return
		}
	}
})
