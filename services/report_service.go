package services

import (
	"kasir-api-bootcamp/models"
	"kasir-api-bootcamp/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodaySummary() (*models.SalesSummary, error) {
	return s.repo.GetTodaySummary()
}

func (s *ReportService) GetSummaryByDateRange(startDate, endDate string) (*models.SalesSummary, error) {
	return s.repo.GetSummaryByDateRange(startDate, endDate)
}
