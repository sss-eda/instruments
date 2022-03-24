package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/stations/{id:[0-9]+}", stationsHandlerFunc).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func stationsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if i, _ := strconv.Atoi(mux.Vars(r)["id"]); i > 10 {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade ws:", err)
			return
		}
		defer c.Close()

		err = c.WriteMessage(websocket.TextMessage, []byte("Hello World! In WS"))
		if err != nil {
			log.Println("write ws:", err)
		}

		return
	}

	w.Write([]byte("Hello World! In HTTP"))
}
