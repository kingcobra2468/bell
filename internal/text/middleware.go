package text

import (
	"time"

	"github.com/go-kit/log"
)

// Middlewere for performing request-based logging of the endpoints.
type LoggingMiddleware struct {
	Logger log.Logger
	Next   TextServicer
}

// Logging wrapper for token registration logic.
func (lm LoggingMiddleware) SendOverSMSGateway(toNumber string, msg string, provider string) error {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"method", "sendoversmsgateway",
			"took", time.Since(begin),
		)
	}(time.Now())

	err := lm.Next.SendOverSMSGateway(toNumber, msg, provider)
	if err != nil {
		lm.Logger.Log("error", err.Error())
	}

	return err
}
