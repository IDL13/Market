package mongodb

import (
	"testing"
	"time"

	"github.com/IDL13/Market/internal/entity"
	"github.com/stretchr/testify/assert"
)

var (
	dt   = time.Date(2020, 12, 12, 0, 0, 0, 0, time.Local)
	from = time.Date(2019, 12, 12, 0, 0, 0, 0, time.Local)
	to   = time.Date(2021, 12, 12, 0, 0, 0, 0, time.Local)
)

func TestAddStatistics(t *testing.T) {
	tests := []struct {
		name      string
		inputBody entity.SaveRequest
		response  int8
	}{{
		name: "OK",
		inputBody: entity.SaveRequest{
			Date:   dt,
			Views:  10,
			Clicks: 10,
			Cost:   14.5,
		},
		response: 1,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mongoDB := New("27017", "test", "test")
			res := mongoDB.AddStatistics(test.inputBody)
			assert.Equal(t, test.response, res)
		})
	}
}

func TestGetStatistics(t *testing.T) {
	tests := []struct {
		name      string
		inputBody entity.ShowRequest
		response  []entity.Statistics
	}{{
		name: "OK",
		inputBody: entity.ShowRequest{
			From: from,
			To:   to,
		},
		response: []entity.Statistics{entity.New(dt, 10, 10, 14.5)},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mongoDB := New("27017", "test", "test")
			_ = mongoDB.AddStatistics(entity.SaveRequest{
				Date:   dt,
				Views:  10,
				Clicks: 10,
				Cost:   14.5,
			})
			res := mongoDB.GetStatistics(test.inputBody)
			assert.NotEqual(t, len(res), 0)
			assert.Equal(t, res[0].Clicks, int32(10))
		})
	}
}

func TestResetStatistics(t *testing.T) {
	tests := []struct {
		name     string
		response int8
	}{{
		name:     "OK",
		response: 1,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mongoDB := New("27017", "test", "test")
			res := mongoDB.ResetStatistics()
			assert.Equal(t, test.response, res)
		})
	}
}
