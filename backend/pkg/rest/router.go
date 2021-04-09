// Package router is the wrapped implementation of the gin goland library.
// Provides customizable logic for core web server functionality.
package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Commonly used variables
const (
	v1 = "api/v1/"
)

// Provides data so that REST routes per service can be registered.
type RouteInfo interface {
	GetRouteGroup() RouteGroup
}

// route contains the data to create a route path within gin router engine.
type Route struct {
	Method       string
	RelativePath string
	Handler      gin.HandlerFunc
}

// routeGroup contains the data to create a grouped routes in gin router engine.
type RouteGroup struct {
	GroupPath string
	Routes    []Route
}

// Router encapsulation so that implementations can be hidden from callers.
type Router struct {
	*gin.Engine
}

// NewRouter will return a point to the wrapped router.
func NewRouter(routeInfos []RouteInfo) (*Router, error) {
	engine := createEngine()
	router := Router{engine}
	router.registerRoutes(routeInfos)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	return &router, nil
}

// createEngine will for now just return the default settings of the gin.engine.
func createEngine() *gin.Engine {
	return gin.Default()
}

// registerRoutes takes in struct RouteInfo and will register all the routes specified along with the
// paths and handlers that are provided.
func (router Router) registerRoutes(routeInfos []RouteInfo) {
	for _, routeInfo := range routeInfos {
		routeGroup := routeInfo.GetRouteGroup()
		newGroup := router.Group(v1 + routeGroup.GroupPath)
		// Assign handle to the new group that was created.
		for _, route := range routeGroup.Routes {
			newGroup.Handle(route.Method, route.RelativePath, route.Handler)
		}
	}
}
