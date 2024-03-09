package utils

import (
	"time"

	"github.com/golang-module/carbon/v2"
)

func CurrentTime() *time.Time {
	now := time.Now()

	return &now
}

func GetPreviousDay(t *time.Time) *time.Time {
	previousDay := carbon.CreateFromStdTime(*t).SubDay()

	result := previousDay.StdTime()

	return &result
}

func GetPreviousMonth(t *time.Time) *time.Time {
	previousMonth := carbon.CreateFromStdTime(*t).SubMonth()

	result := previousMonth.StdTime()

	return &result
}

func GetPreviousYear(t *time.Time) *time.Time {
	previousYear := carbon.CreateFromStdTime(*t).SubYear()

	result := previousYear.StdTime()

	return &result
}

func BeginningOfDay(t *time.Time) *time.Time {
	startDayCarbon := carbon.CreateFromStdTime(*t).StartOfDay()

	result := startDayCarbon.StdTime()

	return &result
}

func EndOfDay(t *time.Time) *time.Time {
	endDayCarbon := carbon.CreateFromStdTime(*t).EndOfDay()

	result := endDayCarbon.StdTime()

	return &result
}

func BeginningOfMonth(t *time.Time) *time.Time {
	startMonth := carbon.CreateFromStdTime(*t).StartOfMonth()

	result := startMonth.StdTime()

	return &result
}

func EndOfMonth(t *time.Time) *time.Time {
	endMonth := carbon.CreateFromStdTime(*t).EndOfMonth()

	result := endMonth.StdTime()

	return &result
}

func BeginningOfYear(t *time.Time) *time.Time {
	startYear := carbon.CreateFromStdTime(*t).StartOfYear()

	result := startYear.StdTime()

	return &result
}

func EndOfYear(t *time.Time) *time.Time {
	endYear := carbon.CreateFromStdTime(*t).EndOfYear()

	result := endYear.StdTime()

	return &result
}
