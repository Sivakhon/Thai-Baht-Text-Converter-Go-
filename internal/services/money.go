package services

import (
	"fmt"
	"strconv"
	"strings"

	gothaibaht "github.com/artykaikub/go-thaibaht"
	"github.com/shopspring/decimal"
	"github.com/sivakhon/float2text/internal/model"
)

var (
	thaiNumWords = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
	unitWords    = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}
)

func FloatToTextPackage(f model.Money) (string, error) {
	bath_string := gothaibaht.ToText(f.Money.InexactFloat64())
	return bath_string, nil
}

func FloatToTextDefault(f model.Money) (string, error) {

	num := f.Money
	zero := decimal.NewFromInt(0)
	result := ""

	// Get integer part and turn to decimal.Decimal
	intPart := num.IntPart() // int64
	bath := decimal.NewFromInt(intPart)

	// Convert to 2-digit integer (e.g., 0.56 → 56)
	decimalPart := num.Sub(bath).Abs()
	satang := decimalPart.Mul(decimal.NewFromInt(100))

	if bath.GreaterThan(zero) {
		result += convertIntegerToThaiText(bath.IntPart()) + "บาท"
	} else {
		result += "ศูนย์บาท"
	}
	if satang.GreaterThan(zero) {
		result += convertIntegerToThaiText(satang.IntPart()) + "สตางค์"
	} else {
		result += "ถ้วน"
	}

	return result, nil
}

func convertIntegerToThaiText(num int64) string { //ยัดเข้าไปทีละล้าน
	if num == 0 {
		return thaiNumWords[0]
	}

	result := ""
	group := 0

	for num > 0 {
		part := num % 1000000
		if part != 0 {
			text := convertBelowMillionToThai(part)
			if group > 0 {
				text += unitWords[6] // "ล้าน"
			}
			result = text + result
		}
		num /= 1000000
		group++
	}

	return result
}

func convertBelowMillionToThai(num int64) string {
	result := ""
	numStr := fmt.Sprintf("%06d", num) // pad to 6 digits
	digits := strings.Split(numStr, "")

	for i := 0; i < 6; i++ {
		d := digits[i]
		n, _ := strconv.Atoi(d)
		if n == 0 {
			continue
		}

		unitIndex := 5 - i

		if unitIndex == 1 && n == 1 {
			result += unitWords[unitIndex]
		} else if unitIndex == 1 && n == 2 {
			result += "ยี่" + unitWords[unitIndex]
		} else if unitIndex == 0 && n == 1 && len(result) > 0 {
			result += "เอ็ด"
		} else {
			result += thaiNumWords[n] + unitWords[unitIndex]
		}
	}

	return result
}
