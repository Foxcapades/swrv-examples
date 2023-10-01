package main

import (
	"fmt"
	"time"

	"github.com/foxcapades/swrv/pkg/swrv"
)

func main() {
	var timeFilter timer
	swrv.NewServer("0.0.0.0", 8080).
		WithRequestFilters(timeFilter).
		WithResponseFilters(timeFilter).
		WithControllers(swrv.NewController("/time", swrv.RequestHandlerFunc(func(_ swrv.Request) swrv.Response {
			return swrv.NewResponse().WithCode(204)
		}))).
		Start(nil)
}

type timer struct{}

func (_ timer) FilterRequest(request swrv.Request) swrv.Response {
	request.AdditionalContext().Put("start-time", time.Now())
	return nil
}

func (_ timer) FilterResponse(request swrv.Request, response swrv.Response) swrv.Response {
	fmt.Printf("request took %s\n", time.Now().Sub(request.AdditionalContext().Get("start-time").(time.Time)))
	return response
}
