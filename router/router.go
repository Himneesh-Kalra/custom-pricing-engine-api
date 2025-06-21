package router

import (
	"net/http"

	"github.com/Himneesh-Kalra/custom-pricing-engine-api/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/compute", handler.ComputeHandler)
	return mux
}
