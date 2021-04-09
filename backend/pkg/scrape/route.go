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

// scrapeData is the json POST body received from UI.
type scrapeData struct {
	URL string `json:"url"`
}

// scrapeWebUrlResp is the json response to send back to UI.
type scrapeWebUrlResp struct {
	BackendLogic string `json:"backend_logic"`
}

// scrapeWebUrl serves POST call to the model layer.
func (sri ScrapeRouteInfo) scrapeWebUrl(c *gin.Context) {
	var data scrapeData
	if err := c.ShouldBindJSON(&data); err == nil {
		// here we are technically calling the implementation through the sc interface.
		processed := sri.sc.ScrapeWebUrl(data.URL)
		resp := scrapeWebUrlResp{BackendLogic: processed}
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
