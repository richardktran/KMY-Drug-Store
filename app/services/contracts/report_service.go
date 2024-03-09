package contracts

import (
	"time"

	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type IReportService interface {
	GetRevenuesByRange(
		from *time.Time,
		to *time.Time,
	) (*models.Revenue, *app.AppError)
	GetYearRevenue(timePoint time.Time) (*models.Revenue, *app.AppError)
	GetMonthRevenue(timePoint time.Time) (*models.Revenue, *app.AppError)
	GetDayRevenue(timePoint time.Time) (*models.Revenue, *app.AppError)
	GetYearRevenueReport(timePoint time.Time) (models.RevenueReport, *app.AppError)
	GetMonthRevenueReport(timePoint time.Time) (models.RevenueReport, *app.AppError)
	GetDayRevenueReport(timePoint time.Time) (models.RevenueReport, *app.AppError)
	GetRevenueReports(timePoint time.Time, rangeType string) (models.RevenueReport, *app.AppError)
}
