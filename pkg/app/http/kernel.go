package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaofan193/xifancloud193/internal/framework"
)

// NewHttpEngine create a web engine whice bind route
func NewHttpEngine(container framework.Container) (*framework.GinEngine, error) {
	// set Release,To ensure that debugging information is not output by default during startup
	gin.SetMode(gin.ReleaseMode)
	// startup a deafault web engine
	gi := &framework.GinEngine{}
	r := gi.New()
	// todo seted the engine
	r.SetContainer(container)
	// defalult regiser recovery middeware
	r.Use(gin.Recovery())
	// todo business bind route operation
	Route(r)
	return r, nil
}
