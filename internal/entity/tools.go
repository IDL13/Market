package entity

import (
	"strconv"
	"strings"
	"time"

	api "github.com/IDL13/Market/pkg/api/proto"
)

func StrToDate(date string) time.Time {
	d := strings.Split(date, "-")

	f := func(s string) int {
		integer, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return integer
	}

	return time.Date(f(d[0]), time.Month(f(d[1])), f(d[2]), 0, 0, 0, 0, time.Local)
}

func StatisticsConvertor(arr []Statistics) []*api.Statistics {
	newArr := make([]*api.Statistics, len(arr))

	for _, statistic := range arr {
		newArr = append(newArr, &api.Statistics{
			Date:   statistic.Date.String(),
			Views:  statistic.Views,
			Clicks: statistic.Clicks,
			Cost:   float32(statistic.Cost),
			Cpc:    float32(statistic.Cpc),
			Cpm:    float32(statistic.Cpm),
		})
	}

	return newArr
}
