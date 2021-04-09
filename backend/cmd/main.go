package main

import (
	"purerecipes/pkg/login"
	rest "purerecipes/pkg/rest"
	"purerecipes/pkg/scrape"
)

func main() {
	// Create managers for all packages.
	lc := login.NewLoginManager()
	lri := login.NewLoginRouteInfo(lc)
	sc := scrape.NewScrapeManager()
	sri := scrape.NewScrapeRouteInfo(sc)

	// Create the routers for all packages.
	routeInfos := []rest.RouteInfo{lri, sri}
	router, _ := rest.NewRouter(routeInfos)
	router.Run(":8000")
}
