package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/sivakhon/float2text/internal/model"
	"github.com/sivakhon/float2text/internal/router"
	"github.com/sivakhon/float2text/internal/services"
)

func main() {
	testMode := flag.Bool("test", false, "Run float-to-text test instead of starting API server")
	flag.Parse()

	if *testMode {
		runTest()
		return
	}

	engine := gin.New()
	router.SetupRouter(engine)

	if err := engine.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func runTest() {
	input := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(21.21),
		decimal.NewFromFloat(21212121.21),
		decimal.NewFromFloat(21212121),
	}

	fmt.Println("Testing FloatToText Default:")
	for _, input := range input {

		model := model.Money{
			Money: input,
		}
		bath_string, _ := services.FloatToTextDefault(model)
		fmt.Println(bath_string)
	}
	fmt.Println("======================================================")
	fmt.Println("Testing FloatToText Package (using gothaibaht):")
	fmt.Println("(this service will be inexact float because it's turning decimal.Decimal to float64)")
	for _, input := range input {
		model := model.Money{
			Money: input,
		}
		bath_string, _ := services.FloatToTextDefault(model)
		fmt.Println(bath_string)
	}
}
