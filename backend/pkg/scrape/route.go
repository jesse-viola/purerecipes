package scrape

import (
	"net/http"
	"purerecipes/pkg/rest"

	"github.com/gin-gonic/gin"
)

// ScrapeCruder for the encapsulation for the ScrapeManager.
type ScrapeCruder interface {
	ScrapeWebUrl(url string) string
}

// ScrapeRouteInfo contains the reference to the ScrapeCruder.
type ScrapeRouteInfo struct {
	sc ScrapeCruder
}

// NewScrapeRouteInfo provides the route for the scrape package.
func NewScrapeRouteInfo(sc ScrapeCruder) ScrapeRouteInfo {
	return ScrapeRouteInfo{sc}
}

// GetRouteGroup ScrapeRouteInfo implementation for package scrape.
func (sri ScrapeRouteInfo) GetRouteGroup() rest.RouteGroup {
	routeGroup := rest.RouteGroup{
		GroupPath: "scrape",
		Routes: []rest.Route{
			{
				Method:       http.MethodPost,
				RelativePath: "",
				Handler:      sri.scrapeWebUrl,
			},
		},
	}
	return routeGroup
}

type scrapeWebUrlResponse struct {
	URL  string `json:"url"`
	Test string
}

func (sri ScrapeRouteInfo) scrapeWebUrl(c *gin.Context) {
	// Call the manager
	var scrapeData scrapeWebUrlResponse
	// scrapeData := sri.sc.ScrapeWebUrl(stringBody) TODO implmentation
	c.BindJSON(&scrapeData)
	response := &scrapeWebUrlResponse{
		URL:  scrapeData.URL,
		Test: "you did it harry", // TODO remove
	}
	c.JSON(http.StatusOK, response)
}
