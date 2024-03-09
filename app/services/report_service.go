package services

import (
	"time"

	"github.com/richardktran/KMY-Drug-Store/app/models"
	repositories "github.com/richardktran/KMY-Drug-Store/app/respositories"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/app/utils"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type ReportService struct {
	orderRepository repositories.OrderRepository
}

func NewReportService(orderRepository repositories.OrderRepository) contracts.IReportService {
	return ReportService{
		orderRepository: orderRepository,
	}
}

func (svc ReportService) GetRevenuesByRange(
	from *time.Time,
	to *time.Time,
) (*models.Revenue, *app.AppError) {
	var revenue models.Revenue
	var err *app.AppError

	currentTotal, err := svc.orderRepository.GetTotalRevenueRange(from, to)

	if err != nil {
		return nil, err
	}

	revenue = models.Revenue{
		From:  from,
		To:    to,
		Total: currentTotal,
	}

	return &revenue, err
}

func (svc ReportService) GetYearRevenue(timePoint time.Time) (*models.Revenue, *app.AppError) {
	var total *models.Revenue
	var err *app.AppError
	var from *time.Time
	var to *time.Time

	now := timePoint

	from = utils.BeginningOfYear(&now)
	to = utils.EndOfYear(&now)
	total, err = svc.GetRevenuesByRange(from, to)
	return total, err
}

func (svc ReportService) GetMonthRevenue(timePoint time.Time) (*models.Revenue, *app.AppError) {
	var total *models.Revenue
	var err *app.AppError
	var from *time.Time
	var to *time.Time

	now := timePoint

	from = utils.BeginningOfMonth(&now)
	to = utils.EndOfMonth(&now)
	total, err = svc.GetRevenuesByRange(from, to)

	return total, err
}

func (svc ReportService) GetDayRevenue(timePoint time.Time) (*models.Revenue, *app.AppError) {
	var total *models.Revenue
	var err *app.AppError
	var from *time.Time
	var to *time.Time

	now := timePoint

	from = utils.BeginningOfDay(&now)
	to = utils.EndOfDay(&now)
	total, err = svc.GetRevenuesByRange(from, to)

	return total, err
}

// Get reports

func (svc ReportService) GetRevenueReports(timePoint time.Time, rangeType string) (models.RevenueReport, *app.AppError) {
	var report models.RevenueReport
	var err *app.AppError

	switch rangeType {
	case "year":
		report, err = svc.GetYearRevenueReport(timePoint)
	case "month":
		report, err = svc.GetMonthRevenueReport(timePoint)
	case "day":
		report, err = svc.GetDayRevenueReport(timePoint)
	default:
		report, err = svc.GetTotalRevenueReport(timePoint)

	}

	return report, err
}

func (svc ReportService) GetYearRevenueReport(timePoint time.Time) (models.RevenueReport, *app.AppError) {
	now := timePoint

	current, err := svc.GetYearRevenue(now)
	if err != nil {
		current = nil
	}

	previousTime := utils.GetPreviousYear(&now)

	previous, err := svc.GetYearRevenue(*previousTime)

	if err != nil {
		previous = nil
	}

	percentageChange := svc.CalculatePercentageChange(previous, current)

	return models.RevenueReport{
		Current:          current,
		Previous:         previous,
		PercentageChange: percentageChange,
	}, nil
}

func (svc ReportService) GetMonthRevenueReport(timePoint time.Time) (models.RevenueReport, *app.AppError) {
	now := timePoint

	current, err := svc.GetMonthRevenue(now)
	if err != nil {
		current = nil
	}

	previousTime := utils.GetPreviousMonth(&now)

	previous, err := svc.GetMonthRevenue(*previousTime)

	if err != nil {
		previous = nil
	}

	percentageChange := svc.CalculatePercentageChange(previous, current)

	return models.RevenueReport{
		Current:          current,
		Previous:         previous,
		PercentageChange: percentageChange,
	}, nil

}

func (svc ReportService) GetDayRevenueReport(timePoint time.Time) (models.RevenueReport, *app.AppError) {
	now := timePoint

	current, err := svc.GetDayRevenue(now)
	if err != nil {
		current = nil
	}

	previousTime := utils.GetPreviousDay(&now)

	previous, err := svc.GetDayRevenue(*previousTime)

	if err != nil {
		previous = nil
	}

	percentageChange := svc.CalculatePercentageChange(previous, current)

	return models.RevenueReport{
		Current:          current,
		Previous:         previous,
		PercentageChange: percentageChange,
	}, nil
}

func (svc ReportService) GetTotalRevenueReport(timePoint time.Time) (models.RevenueReport, *app.AppError) {
	now := timePoint

	current, err := svc.GetRevenuesByRange(nil, nil)
	if err != nil {
		current = nil
	}

	previousTime := utils.GetPreviousDay(&now)

	previous, err := svc.GetRevenuesByRange(previousTime, previousTime)

	if err != nil {
		previous = nil
	}

	percentageChange := 100.0

	return models.RevenueReport{
		Current:          current,
		Previous:         previous,
		PercentageChange: percentageChange,
	}, nil
}

func (svc ReportService) CalculatePercentageChange(previous *models.Revenue, current *models.Revenue) float64 {
	var percentageChange float64

	if previous != nil && current != nil {
		percentageChange = utils.CalculatePercentageChange(previous.Total, current.Total)
	} else {
		if previous == nil && current == nil {
			percentageChange = 0
		} else if current == nil {
			percentageChange = -100
		} else {
			percentageChange = 100
		}
	}

	return percentageChange
}
