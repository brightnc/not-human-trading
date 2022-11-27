package protocol

import (
	"log"
	"os"
	"os/signal"

	"github.com/brightnc/not-human-trading/config"
	"github.com/brightnc/not-human-trading/internal/handler/httphdl"
	"github.com/brightnc/not-human-trading/internal/handler/ws"
	"github.com/brightnc/not-human-trading/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

/*
	|--------------------------------------------------------------------------
	| Application Protocol
	|--------------------------------------------------------------------------
	|
	| Here you can choose which protocol your application wants to interact
	| with the client for instance HTTP, gRPC etc.
	|
*/

// The example to serve REST
func ServeREST() error {
	srv := fiber.New(fiber.Config{
		DisableKeepalive:      false,
		DisableStartupMessage: true,
	})

	srv.Use(cors.New(cors.ConfigDefault))
	web.Web(srv)
	hdl := httphdl.NewHTTPHandler(app.svc, app.pkg.vld)
	wshdl := ws.NewWebSocketHandler(app.svc, app.pkg.vld)
	v1Group := srv.Group("/v1")
	wsGroup := srv.Group("/ws")
	wsV1Group := wsGroup.Group("/v1")
	wsV1Group.Use("", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	exchangesV1Group := v1Group.Group("/exchanges")
	exchangesV1Group.Put("", hdl.UpdateBotExchangeConfig)
	wsV1Group.Get("", websocket.New(func(c *websocket.Conn) {
		wshdl.SubscribeMessage(c)
	}))
	botsV1Group := v1Group.Group("/bots")
	// example
	botsV1Group.Post("/start", hdl.StartBot)
	botsV1Group.Post("/stop", hdl.StopBot)

	// gracefully shuts down  ...
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("Gracefully shutting down ...")
		err := srv.Shutdown()
		if err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}()
	err := srv.Listen(":" + config.GetConfig().App.HTTPPort)
	if err != nil {
		return err
	}
	return nil
}
