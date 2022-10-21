package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	ws2 "github.com/open-okapi/okapi-server/inner/ws"
	"log"
	"net/http"
)

func setWSRouter(r *gin.Engine) {
	r.GET("/ws", WSHandler)
}

func WSHandler(c *gin.Context) {
	// 升级 get 请求为 webSocket 协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get ws connection")
		log.Println(err)
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Println("ws close error")
			log.Println(err)
		}
	}(ws)
	for {
		var proxy = ws2.ProxyRequest{}
		err = ws.ReadJSON(&proxy)
		if err != nil {
			log.Println("error read json")
			log.Println(err)
			ws2.ErrorResponse(ws, err.Error())
			return
		}
		// router sub method list
		switch proxy.Type {
		case ws2.HELLO:
			ws2.Welcome(ws, proxy)
		}
	}
}

var upGrader = websocket.Upgrader{
	HandshakeTimeout: 0,
	ReadBufferSize:   0,
	WriteBufferSize:  0,
	WriteBufferPool:  nil,
	Subprotocols:     nil,
	Error:            nil,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: false,
}
