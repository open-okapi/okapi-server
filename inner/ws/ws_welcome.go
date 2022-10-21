package ws

import (
	"github.com/gorilla/websocket"
	"github.com/open-okapi/okapi-server/inner/pkg/utils"
	"github.com/open-okapi/okapi-server/inner/ws/model"
)

func Welcome(ws *websocket.Conn, data ProxyRequest) {
	SuccessResponse(ws, data, utils.Obj2Str(model.Welcome{
		Say: "welcome " + data.Data,
	}))
}
