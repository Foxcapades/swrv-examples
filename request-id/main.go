package main

import (
	"github.com/foxcapades/swrv/pkg/swrv"
	"github.com/teris-io/shortid"
)

type Response struct {
	RequestID string `json:"request-id"`
	Message   string `json:"message"`
}

func requestIDFilter(request swrv.Request) swrv.Response {
	request.AdditionalContext().Put("request-id", shortid.MustGenerate())
	return nil
}

func handler(request swrv.Request) swrv.Response {
	return swrv.NewResponse().
		WithBody(Response{
			RequestID: request.AdditionalContext().Get("request-id").(string),
			Message:   "I generated a request ID!",
		})
}

func main() {
	swrv.NewServer("0.0.0.0", 8080).
		WithRequestFilters(swrv.RequestFilterFunc(requestIDFilter)).
		WithObjectSerializers(swrv.NewDefaultJSONObjectSerializer()).
		WithControllers(swrv.NewController("/request-id", swrv.RequestHandlerFunc(handler))).
		Start(nil)
}
