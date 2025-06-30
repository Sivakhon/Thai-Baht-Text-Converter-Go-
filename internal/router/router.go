package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sivakhon/float2text/internal/handler"
)

func SetupRouter(r *gin.Engine) {
	// Set up the router with the handler
	r.GET("/moneydefault", handler.MoneyDefault)
	r.GET("/moneygothaibaht", handler.MoneyGothaiBaht)
}
