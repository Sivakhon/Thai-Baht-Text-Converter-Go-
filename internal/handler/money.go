package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtomapper "github.com/sivakhon/float2text/internal/model/dto-mapper"
	"github.com/sivakhon/float2text/internal/services"

	"github.com/sivakhon/float2text/internal/pkg/dto"
)

func MoneyDefault(c *gin.Context) {
	var (
		req dto.Money
		err error
	)

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	resp, err := services.FloatToTextDefault(dtomapper.MoneyRquestToModel(req))
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	response := dto.MoneyResponse{
		Text: resp,
	}
	c.JSON(http.StatusOK, response)
}

func MoneyGothaiBaht(c *gin.Context) {
	var (
		req dto.Money
		err error
	)

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	resp, err := services.FloatToTextPackage(dtomapper.MoneyRquestToModel(req))
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	response := dto.MoneyResponse{
		Text: resp,
	}
	c.JSON(http.StatusOK, response)
}
