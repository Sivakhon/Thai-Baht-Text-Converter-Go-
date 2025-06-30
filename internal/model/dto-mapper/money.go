package dtomapper

import (
	"github.com/shopspring/decimal"
	"github.com/sivakhon/float2text/internal/model"
	"github.com/sivakhon/float2text/internal/pkg/dto"
)

func MoneyRquestToModel(req dto.Money) model.Money {
	return model.Money{
		Money: decimal.NewFromFloat(req.Money)}
}
