package protocol

import (
	"github.com/brightnc/not-human-trading/config"
	"github.com/brightnc/not-human-trading/internal/core/service"
	"github.com/brightnc/not-human-trading/internal/repository"
	"github.com/brightnc/not-human-trading/pkg/logger"
	"github.com/brightnc/not-human-trading/pkg/validators"
)

var app *application

type application struct {
	// svc stand for service
	svc *service.Service
	// pkg stand for package
	pkg packages
}

type packages struct {
	vld validators.Validator
}

func init() {
	logger.Init("BOT-XXX")
	config.Init()
	biananceRepo := repository.NewBinanceExchange()
	botConfigRepo := repository.NewBotConfig()
	packages := packages{
		vld: validators.New(),
	}
	//todo: inject repository into the service
	app = &application{
		svc: service.NewService(biananceRepo, botConfigRepo),
		pkg: packages,
	}
}
