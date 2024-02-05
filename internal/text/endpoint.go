package text

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type sendOverSMSRequest struct {
	ToNumber string `json:"to_number"`
	Message  string `json:"message"`
	Provider string `json:"provider"`
}

type sendOverSMSResponse struct {
	Status  string  `json:"status"`
	Data    *string `json:"data"`
	Message string  `json:"message,omitempty"`
}

func makeSendOverSMSGateway(ts TextServicer) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(sendOverSMSRequest)

		err := ts.SendOverSMSGateway(req.ToNumber, req.Message, req.Provider)
		if err != nil {
			return sendOverSMSResponse{Status: "error", Message: err.Error()}, nil
		}

		return sendOverSMSResponse{Status: "success"}, nil
	}
}
