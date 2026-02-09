package handlers

import (
	"net/http"
	"strings"

	resp "kasir-api-bootcamp/common/handlers"
	"kasir-api-bootcamp/services"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetReport(w, r)
	default:
		resp.WriteJSON(w, http.StatusMethodNotAllowed, resp.Response{
			Status:  "error",
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
			Data:    nil,
		})
	}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/report")

	if path == "/today" || path == "" && r.URL.Query().Get("start_date") == "" {
		summary, err := h.service.GetTodaySummary()
		if err != nil {
			resp.WriteJSON(w, http.StatusInternalServerError, resp.Response{
				Status:  "error",
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		resp.WriteJSON(w, http.StatusOK, resp.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Today's sales summary",
			Data:    summary,
		})
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "start_date and end_date are required",
			Data:    nil,
		})
		return
	}

	summary, err := h.service.GetSummaryByDateRange(startDate, endDate)
	if err != nil {
		resp.WriteJSON(w, http.StatusInternalServerError, resp.Response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, resp.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Sales summary",
		Data:    summary,
	})
}
