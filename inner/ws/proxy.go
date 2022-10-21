package ws

import (
	"github.com/gorilla/websocket"
	"log"
)

type ProxyRequest struct {
	UUID string    `json:"uuid"`
	Type ProxyType `json:"type"`
	Data string    `json:"data"`
}

type ProxyResponse struct {
	UUID    string    `json:"uuid"`
	Type    ProxyType `json:"type"`
	Data    string    `json:"data"`
	Status  int32     `json:"status"`
	Message string    `json:"message"`
}

type ProxyType string

const (
	HELLO ProxyType = "hello"
)

func SuccessResponse(ws *websocket.Conn, pr ProxyRequest, dataStr string) {
	err := ws.WriteJSON(ProxyResponse{
		UUID:    pr.UUID,
		Type:    pr.Type,
		Status:  200,
		Message: "success",
		Data:    dataStr,
	})
	if err != nil {
		log.Println("error write response success: " + err.Error())
	}
}

func WarnResponse(ws *websocket.Conn, pr ProxyRequest, messageStr string, dataStr string) {
	err := ws.WriteJSON(ProxyResponse{
		UUID:    pr.UUID,
		Type:    pr.Type,
		Status:  -2,
		Message: messageStr,
		Data:    dataStr,
	})
	if err != nil {
		log.Println("error write response warn: " + err.Error())
	}
}

func ErrorResponse(ws *websocket.Conn, errStr string) {
	err := ws.WriteJSON(ProxyResponse{
		UUID:    "",
		Type:    "",
		Status:  -1,
		Message: errStr,
		Data:    "",
	})
	if err != nil {
		log.Println("error write response error: " + err.Error())
	}
}
