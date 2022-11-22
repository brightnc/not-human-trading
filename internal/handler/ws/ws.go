package ws

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brightnc/not-human-trading/internal/core/domain"
	"github.com/brightnc/not-human-trading/internal/core/port"
	"github.com/brightnc/not-human-trading/pkg/validators"
	"github.com/gofiber/websocket/v2"
)

type WebSocketHandler struct {
	svc       port.Service
	validator validators.Validator
}

type subscription struct {
	connection *domain.Connection
}

func NewWebSocketHandler(svc port.Service, validators validators.Validator) *WebSocketHandler {
	return &WebSocketHandler{
		svc:       svc,
		validator: validators,
	}
}

func (r *WebSocketHandler) SubscribeMessage(c *websocket.Conn) {
	sub := &subscription{
		connection: &domain.Connection{Ws: c, Message: make(chan domain.WsMessage)},
	}
	r.svc.RegisterWebsocketClient(sub.connection)
	sub.writeMessage()

}

func (r *subscription) writeMessage() {
	c := r.connection
	defer func() {
		r.connection.Ws.Close()
	}()

	for {
		select {
		case message := <-c.Message:
			fmt.Println("Broadcast Message  ", message)
			_, err := json.Marshal(message)
			if err != nil {
				fmt.Println("CANNOT MARSHAL JSON ", err)
				panic(err)
			}
			if c.Ws == nil {
				fmt.Println("Nil Connection")
			}
			c.Ws.SetWriteDeadline(time.Now().Add(time.Second * 4))
			err = c.Ws.WriteJSON(message)
			if err != nil {
				fmt.Println("CANNOT WRITE JSON ", err)
				panic(err)
			}
		}
	}
}
