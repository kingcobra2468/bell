package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"

	"github.com/caarlos0/env"
	"github.com/kingcobra2468/bell/internal/text"
)

type config struct {
	ServiceHostname string `env:"BELL_HOSTNAME" envDefault:"127.0.0.1"`
	ServicePort     int    `env:"BELL_PORT" envDefault:"8080"`
	Email           string `env:"BELL_EMAIL"`
	Password        string `env:"BELL_EMAIL_PWD"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	auth := smtp.PlainAuth("", cfg.Email, cfg.Password, "smtp.gmail.com")
	var service text.TextServicer = text.Text{Auth: auth}
	service = text.LoggingMiddleware{Logger: logger, Next: service}
	var h http.Handler = text.MakeHTTPHandler(service)

	errs := make(chan error)
	// Listener for Ctrl+C signals
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	// Launch microservice
	go func() {
		url := fmt.Sprintf("%s:%d", cfg.ServiceHostname, cfg.ServicePort)

		logger.Log("transport", "HTTP", "addr", url)
		errs <- http.ListenAndServe(url, h)
	}()

	logger.Log("exit", <-errs)
}
