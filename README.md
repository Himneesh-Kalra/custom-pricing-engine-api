# Custom Pricing Engine API

A simple Go-based RESTful API that simulates a pricing engine. It accepts a list of product orders, applies discount rules, and returns the calculated pricing breakdown for each product along with the overall total.

---

## Features

* RESTful API (`POST /compute`)
* Calculates price per product
* Applies flat and percentage discounts
* Returns final prices with discount breakdown
* Includes unit and integration tests

---

## Prerequisites

* Go 1.18 or above
* Git 
* cURL or Postman (for testing)

---

## How to Run

### 1. Clone the repository

```bash
git clone https://github.com/Himneesh-Kalra/custom-pricing-engine-api.git
cd custom-pricing-engine-api
```

### 2. Run the server




```bash
go run main.go
```

---

## Sample Payloads for Testing

### Payload 1 (Basic discounts)

```json
{
  "orders": [
    { "product_id": "A101", "quantity": 3, "unit_price": 100 },
    { "product_id": "C303", "quantity": 1, "unit_price": 500 }
  ],
  "discount_rules": [
    { "product_id": "A101", "type": "percentage", "value": 10 },
    { "product_id": "C303", "type": "flat", "value": 50 }
  ]
}
```

### Payload 2 (Multiple products, mixed discounts)

```json
{
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
```

### Payload 3 (Invalid JSON to test error handling)

```json
{ invalid json }
```

---

## Running Tests

To run all unit and integration tests:

```bash
go test ./...
```

To run tests with verbose output:

```bash
go test ./... -v
```

---

## Project Structure

```
.
├── handler               # HTTP handlers
├── logic                 # Pricing logic
├── models                # Data models
├── router                # HTTP routes
├── main.go               # Application entrypoint
├── go.mod / go.sum       # Go dependencies
├── README.md             # Project documentation
```

---

