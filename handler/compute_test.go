package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestComputeHandler(t *testing.T) {
	tests := []struct {
		name              string
		reqBody           string
		expectedStatus    int
		expectedSubstring string
	}{
		{
			name: "Valid json payload",
			reqBody: `{
						"orders": [
							{ "product_id": "A101", "quantity": 3, "unit_price": 100 },
							{ "product_id": "C303", "quantity": 1, "unit_price": 500 }
						],
						"discount_rules": [
							{ "product_id": "A101", "type": "percentage", "value": 10 },
							{ "product_id": "C303", "type": "flat", "value": 50 }
						]
					}`,
			expectedStatus:    http.StatusOK,
			expectedSubstring: `"total": 720`,
		}, {
			name: "Valid json payload 2",
			reqBody: `{
						"orders": [
							{ "product_id": "PHX01", "quantity": 2, "unit_price": 450 },
							{ "product_id": "DRK99", "quantity": 5, "unit_price": 120 },
							{ "product_id": "LGT88", "quantity": 1, "unit_price": 800 }
						],
						"discount_rules": [
							{ "product_id": "PHX01", "type": "flat", "value": 100 },
							{ "product_id": "DRK99", "type": "percentage", "value": 20 }
						]
					}
					`,
			expectedStatus:    http.StatusOK,
			expectedSubstring: `"total": 2080`,
		}, {
			name:              "Invalid json payload",
			reqBody:           `{something something random payload}`,
			expectedStatus:    http.StatusBadRequest,
			expectedSubstring: "Invalid JSON",
		},
	}

	for _, tt := range tests {

		req := httptest.NewRequest("POST", "/compute", strings.NewReader(tt.reqBody))
		req.Header.Set("Content-Type", "application/json")

		respRec := httptest.NewRecorder()

		ComputeHandler(respRec, req)

		if respRec.Code != tt.expectedStatus {
			t.Errorf("expected status 200, got %+v ", respRec.Code)
		}

		expectedSubstring := tt.expectedSubstring

		if !strings.Contains(respRec.Body.String(), expectedSubstring) {
			t.Errorf("expected resp body to contain %v , but got %v ", expectedSubstring, respRec.Body.String())
		}
	}
}
