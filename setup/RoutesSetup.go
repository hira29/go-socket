package setup

import (
	"go-socket/controller"
	"net/http"
)


func SetupRoutes(){
	http.HandleFunc("/", controller.WsEndPoint)
	http.HandleFunc("/ws", controller.WsEndPoint)
}
