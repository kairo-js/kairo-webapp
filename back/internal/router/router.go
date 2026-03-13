package router

import (
	"net/http"

	"github.com/shizuku86/kairo-webapp/back/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/health", handler.Health)

	return mux
}