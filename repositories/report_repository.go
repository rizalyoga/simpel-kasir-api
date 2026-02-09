package repositories

import (
	"database/sql"

	"kasir-api-bootcamp/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetTodaySummary() (*models.SalesSummary, error) {
	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`

	var summary models.SalesSummary
	err := repo.db.QueryRow(query).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	topProductQuery := `
		SELECT 
			p.name,
			COALESCE(SUM(td.quantity), 0) as qty_terjual
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`

	var topProduct models.ProdukTerlaris
	err = repo.db.QueryRow(topProductQuery).Scan(&topProduct.Nama, &topProduct.QtyTerjual)
	if err == sql.ErrNoRows {
		topProduct = models.ProdukTerlaris{Nama: "-", QtyTerjual: 0}
	} else if err != nil {
		return nil, err
	}

	summary.ProdukTerlaris = topProduct
	return &summary, nil
}

func (repo *ReportRepository) GetSummaryByDateRange(startDate, endDate string) (*models.SalesSummary, error) {
	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE DATE(created_at) BETWEEN $1 AND $2
	`

	var summary models.SalesSummary
	err := repo.db.QueryRow(query, startDate, endDate).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	topProductQuery := `
		SELECT 
			p.name,
			COALESCE(SUM(td.quantity), 0) as qty_terjual
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE DATE(t.created_at) BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`

	var topProduct models.ProdukTerlaris
	err = repo.db.QueryRow(topProductQuery, startDate, endDate).Scan(&topProduct.Nama, &topProduct.QtyTerjual)
	if err == sql.ErrNoRows {
		topProduct = models.ProdukTerlaris{Nama: "-", QtyTerjual: 0}
	} else if err != nil {
		return nil, err
	}

	summary.ProdukTerlaris = topProduct
	return &summary, nil
}
