package entity

import (
	"math"
	"time"
)

type Statistics struct {
	Date   time.Time `json:"date"`
	Views  int32     `json:"views"`
	Clicks int32     `json:"cliks"`
	Cost   float64   `json:"cost"`
	Cpc    float64   `json:"cpc"`
	Cpm    float64   `json:"cmp"`
}

func New(data time.Time, views, clicks int32, cost float64) Statistics {
	return Statistics{
		Date:   data,
		Views:  views,
		Clicks: clicks,
		Cost:   cost,
		Cpc:    math.Floor(cost / float64(clicks)),
		Cpm:    math.Floor(cost / float64(views) * 1000),
	}
}

type SaveReuqestAdapter struct {
	Date   string  `json:"date"`
	Views  int32   `json:"views"`
	Clicks int32   `json:"clicks"`
	Cost   float32 `json:"cost"`
}

type SaveRequest struct {
	Date   time.Time `json:"date"`
	Views  int32     `json:"views"`
	Clicks int32     `json:"clicks"`
	Cost   float32   `json:"cost"`
}

func NewSaveRequest(sra SaveReuqestAdapter) SaveRequest {
	return SaveRequest{
		Date:   StrToDate(sra.Date),
		Views:  sra.Views,
		Clicks: sra.Clicks,
		Cost:   sra.Cost,
	}
}

type ShowRequest struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func NewShowRequest(sra ShowRequestAdapter) ShowRequest {
	return ShowRequest{
		From: StrToDate(sra.From),
		To:   StrToDate(sra.To),
	}
}

type ShowRequestAdapter struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type StatisticsResponse struct {
	Statistics []Statistics `json:"statistics"`
}
