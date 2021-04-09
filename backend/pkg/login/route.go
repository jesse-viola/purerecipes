package login

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"purerecipes/pkg/rest"
)

// LoginCruder for the encapsulation between LoginManager.
type LoginCruder interface {
	Get()
}

// LoginRouteInfo contains reference to the LoginCruder.
type LoginRouteInfo struct {
	lc LoginCruder
}

// NewLoginRouteInfo provides the route for the login package.
func NewLoginRouteInfo(lc LoginCruder) LoginRouteInfo {
	return LoginRouteInfo{lc}
}

// GetRouteGroup LoginRouteInfo implementation for package login.
func (lri LoginRouteInfo) GetRouteGroup() rest.RouteGroup {
	routeGroup := rest.RouteGroup{
		GroupPath: "login",
		Routes: []rest.Route{
			{
				Method:       http.MethodGet,
				RelativePath: "",
				Handler:      lri.get,
			},
		},
	}
	return routeGroup
}

// get handler function for the rest.RouteGroup.
func (lri LoginRouteInfo) get(c *gin.Context) {
	lri.lc.Get()
	c.JSON(http.StatusOK, "Logged in! (not really)")
}
