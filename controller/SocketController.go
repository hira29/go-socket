package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	iter := 1
	for {
		p := make(map[string]interface{})
		now := time.Now()
		now.Format(time.RFC3339)
		if iter > 1 {
			addDate := iter - 1
			now = now.AddDate(0, 0 , addDate)
		}
		datetime := now.Format("2006-01-02")
		p["date"] = datetime
		p["count"] = rand.Intn(100)
		fmt.Println(p)

		data, _ := json.Marshal(p)
		messageType := 1
		if err := conn.WriteMessage(messageType, []byte(data)); err != nil {
			log.Println(err)
			return
		}
		iter += 1;
		time.Sleep(2 * time.Second)
	}
}

func WsEndPoint(w http.ResponseWriter, r*http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Successfully Connected...")

	reader(ws)
}