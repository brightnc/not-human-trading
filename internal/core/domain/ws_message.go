package domain

import (
	"time"

	"github.com/gofiber/websocket/v2"
)

type WsMessage struct {
	Time    time.Time   `json:"time"`
	Message string      `json:"message"`
	Type    messageType `json:"type"`
	SubID   int         `json:"sub_id"`
}

type messageType string

const (
	MessageTypeFeed          messageType = "FEED"
	MessageTypeTradingReport messageType = "TRADING_REPORT"
)

type Connection struct {
	Ws      *websocket.Conn
	Message chan WsMessage
}
