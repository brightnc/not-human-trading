package protocol

import (
	"log"
	"os"
	"os/signal"

	"github.com/brightnc/not-human-trading/config"
	"github.com/brightnc/not-human-trading/internal/handler/httphdl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
		DisableKeepalive: false,
	})

	srv.Use(cors.ConfigDefault)
	hdl := httphdl.NewHTTPHandler(app.svc, app.pkg.vld)
	// example
	srv.Put("/indicators", hdl.UpdateIndicator)

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
