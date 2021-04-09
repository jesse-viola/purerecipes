package main

import (
	"purerecipes/pkg/login"
	rest "purerecipes/pkg/rest"
)

func main() {
	// Create managers for all packages.
	lc := login.NewLoginManager()
	lri := login.NewLoginRouteInfo(lc)

	// Create the routers for all packages.
	routeInfos := []rest.RouteInfo{lri}
	router, _ := rest.NewRouter(routeInfos)
	router.Run(":8000")
}
