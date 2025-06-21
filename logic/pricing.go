package logic

import (
	"fmt"

	"github.com/Himneesh-Kalra/custom-pricing-engine-api/models"
)

func CalculatePrice(order models.Order, ruleMap map[string]models.DiscountRule) models.ComputeResult {
	originalPrice := float64(order.Quantity) * *order.UnitPrice
	discount := 0.0

	rule, exists := ruleMap[order.ProductID]
	if exists {
		switch rule.Type {
		case "percentage":
			discount = originalPrice * rule.Value / 100
		case "flat":
			discount = rule.Value
		}

		if discount > originalPrice {
			discount = originalPrice
		}
	}
	finalPrice := originalPrice - discount

	fmt.Println("discount", discount)
	return models.ComputeResult{
		ProductID:       order.ProductID,
		OriginalPrice:   originalPrice,
		DiscountApplied: discount,
		FinalPrice:      finalPrice,
	}
}
