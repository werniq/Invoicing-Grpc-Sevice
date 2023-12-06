package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (app *Application) Serve() error {
	r := gin.Default()

	app.SetUpRoutes(r)

	return r.Run(fmt.Sprintf(
		":%d", app.config.InvoiceHttpServerPort))
}
