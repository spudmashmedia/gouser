package health

import (
	"net/http"

	"github.com/spudmashmedia/gouser/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

type HealthResponse struct {
	Status string `json:"status"`
}

func (h *handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}
	json.Write(w, http.StatusOK, response)
}
