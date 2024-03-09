package controllers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/app/services/contracts"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

type ReportController struct {
	reportService contracts.IReportService
}

func NewReportController(reportService contracts.IReportService) ReportController {
	return ReportController{
		reportService: reportService,
	}
}

func (ctl *ReportController) GetRevenues() func(*gin.Context) {
	return func(c *gin.Context) {
		timePoint := time.Now()
		yearRevenues, err := ctl.reportService.GetRevenueReports(timePoint, "year")

		if err != nil {
			log.Println(err)
		}

		monthRevenues, err := ctl.reportService.GetRevenueReports(timePoint, "month")

		if err != nil {
			log.Println(err)

		}

		dayRevenues, err := ctl.reportService.GetRevenueReports(timePoint, "day")

		if err != nil {
			log.Println(err)
		}

		totalRevenue, err := ctl.reportService.GetRevenueReports(timePoint, "total")

		if err != nil {
			log.Println(err)
		}

		var revenues = map[string]interface{}{
			"year_revenue":  yearRevenues,
			"month_revenue": monthRevenues,
			"day_revenue":   dayRevenues,
			"total_revenue": totalRevenue,
		}

		app.ResponseSuccess(revenues).Context(c)
	}
}
