package service

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

func TestSaveStatistics(t *testing.T) {
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
			mongoDB := NewTestService()
			res := mongoDB.SaveStatistics(test.inputBody)
			assert.Equal(t, test.response, res)
		})
	}
}

func TestGetStatistics(t *testing.T) {
	tests := []struct {
		name      string
		inputBody entity.ShowRequest
	}{{
		name: "OK",
		inputBody: entity.ShowRequest{
			From: from,
			To:   to,
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewTestService()
			res := service.ShowStatistics(test.inputBody)
			assert.Equal(t, len(res), 1)
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
			service := NewTestService()
			res := service.ResetStatistics()
			assert.Equal(t, test.response, res)
		})
	}
}
