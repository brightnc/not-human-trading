package protocol

import (
	"log"
	"os"
	"os/signal"

	"github.com/brightnc/not-human-trading/config"
	"github.com/brightnc/not-human-trading/internal/core/service"
	"github.com/brightnc/not-human-trading/internal/handler/httphdl"
	"github.com/brightnc/not-human-trading/internal/repository"
	"github.com/brightnc/not-human-trading/pkg/validators"

	"github.com/gofiber/fiber/v2"
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
	app := fiber.New(fiber.Config{
		DisableKeepalive: false,
	})

	biananceRepo := repository.NewBinanceExchange()
	indicatorRepo := repository.NewIndicator()
	svc := service.NewService(biananceRepo, indicatorRepo)
	vld := validators.New()
	hdl := httphdl.NewHTTPHandler(svc, vld)

	// example
	app.Put("/indicators", hdl.UpdateIndicator)

	// gracefully shuts down  ...
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("Gracefully shutting down ...")
		err := app.Shutdown()
		if err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}()
	err := app.Listen(":" + config.GetConfig().App.HTTPPort)
	if err != nil {
		return err
	}
	return nil
}
