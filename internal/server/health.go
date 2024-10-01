package server

import (
	"net/http"

	"oss/internal/service"
)

func HealthHandler(svc *service.OssService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/pub/healthcheck", svc.HealthCheck)
	return mux
}
