package models

type Order struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type DiscountRule struct {
	ProductID string  `json:"product_id"`
	Type      string  `json:"type"`
	Value     float64 `json:"value"`
}

type ComputeRequest struct {
	Orders        []Order        `json:"orders"`
	DiscountRules []DiscountRule `json:"discount_rules"`
}

type ComputeResult struct {
	ProductID       string  `json:"product_id"`
	OriginalPrice   float64 `json:"original_price"`
	DiscountApplied float64 `json:"discount_applied"`
	FinalPrice      float64 `json:"final_price"`
}

type ComputeResponse struct {
	Results []ComputeResult `json:"results"`
	Total   float64         `json:"total"`
}
