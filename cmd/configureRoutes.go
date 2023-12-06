package main

import "github.com/gin-gonic/gin"

func (app *Application) SetUpRoutes(r *gin.Engine) {
	r.POST("/invoice/generate", app.GenerateInvoice)
}
