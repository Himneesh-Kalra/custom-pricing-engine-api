package logic

import (
	"testing"

	"github.com/Himneesh-Kalra/custom-pricing-engine-api/models"
)

func floatPtr(f float64) *float64 {
	return &f
}

func TestCalculatePrice(t *testing.T) {
	tests := []struct {
		name     string
		order    models.Order
		rules    map[string]models.DiscountRule
		expected models.ComputeResult
	}{
		{
			name: "Percentage discount applied",
			order: models.Order{
				ProductID: "H1012",
				Quantity:  2,
				UnitPrice: floatPtr(200),
			},
			rules: map[string]models.DiscountRule{
				"H1012": {ProductID: "H1012", Type: "percentage", Value: 20},
			},
			expected: models.ComputeResult{
				ProductID:       "H1012",
				OriginalPrice:   400,
				DiscountApplied: 80,
				FinalPrice:      320,
			},
		},
		{
			name: "Flat discount applied",
			order: models.Order{
				ProductID: "PH03N1X",
				Quantity:  4,
				UnitPrice: floatPtr(300),
			},
			rules: map[string]models.DiscountRule{
				"PH03N1X": {ProductID: "PH03N1X", Type: "flat", Value: 70},
			},
			expected: models.ComputeResult{
				ProductID:       "PH03N1X",
				OriginalPrice:   1200,
				DiscountApplied: 70,
				FinalPrice:      1130,
			},
		},
		{
			name: "No discount applied",
			order: models.Order{
				ProductID: "K132",
				Quantity:  5,
				UnitPrice: floatPtr(150),
			},
			rules: map[string]models.DiscountRule{},
			expected: models.ComputeResult{
				ProductID:       "K132",
				OriginalPrice:   750,
				DiscountApplied: 0,
				FinalPrice:      750,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePrice(tt.order, tt.rules)

			if result != tt.expected {
				t.Errorf("got %+v, expected %+v ", result, tt.expected)
			}
		})
	}
}
