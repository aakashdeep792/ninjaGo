package handlers

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service is Up and Running"))
	w.WriteHeader(http.StatusOK)
}
