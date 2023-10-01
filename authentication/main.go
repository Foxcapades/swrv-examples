package main

import (
	"net/http"

	"github.com/foxcapades/swrv/pkg/swrv"
)

func main() {
	swrv.NewServer("0.0.0.0", 8080).
		WithControllers(
			swrv.NewController("/do-something", swrv.RequestHandlerFunc(handler)).
				ForMethods(http.MethodPost, http.MethodPatch),
		).
		WithRequestFilters(swrv.RequestFilterFunc(authFilter)).
		WithObjectSerializers(swrv.NewDefaultJSONObjectSerializer()).
		Start(nil)
}

type Response401 struct {
	Message string `json:"message"`
}

func handler(_ swrv.Request) swrv.Response {
	return swrv.NewResponse().
		WithCode(http.StatusNoContent)
}

func authFilter(request swrv.Request) swrv.Response {
	if request.GetHeader("Auth-Token") != "some-secret-key" {
		return swrv.NewResponse().
			WithCode(401).
			WithBody(Response401{"request unauthorized"})
	}

	return nil
}
