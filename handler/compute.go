package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/Himneesh-Kalra/custom-pricing-engine-api/logic"
	"github.com/Himneesh-Kalra/custom-pricing-engine-api/models"
)

func ComputeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong request method", http.StatusMethodNotAllowed)
		return
	}

	var req models.ComputeRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if len(req.Orders) == 0 {
		http.Error(w, "No orders provided", http.StatusBadRequest)
	}

	ruleMap := make(map[string]models.DiscountRule)
	log.Println(req.DiscountRules)
	for _, rule := range req.DiscountRules {
		ruleMap[rule.ProductID] = rule
	}

	wg := sync.WaitGroup{}

	results := make([]models.ComputeResult, len(req.Orders))
	totalChan := make(chan float64, len(req.Orders))
	log.Println("Rulemap", ruleMap)
	for i, order := range req.Orders {
		wg.Add(1)
		go func(i int, order models.Order) {
			defer wg.Done()

			log.Println("product id", order.ProductID)
			res := logic.CalculatePrice(order, ruleMap)
			results[i] = res
			totalChan <- res.FinalPrice
		}(i, order)
	}
	wg.Wait()
	close(totalChan)

	total := 0.0
	for val := range totalChan {
		total += val
	}

	resp := models.ComputeResponse{
		Results: results,
		Total:   total,
	}

	jsonBytes, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
