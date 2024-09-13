package service

import (
	"github.com/IDL13/Market/internal/adapter/mongodb"
	"github.com/IDL13/Market/internal/entity"
)

type Service interface {
	ResetStatistics() int8
	SaveStatistics(json entity.SaveRequest) int8
	ShowStatistics(json entity.ShowRequest) []entity.Statistics
}

type service struct {
	mongodb mongodb.MongoDB
}

func NewService() Service {
	return service{
		mongodb: mongodb.New("27017", "statistics", "statistics"),
	}
}

func NewTestService() Service {
	return service{
		mongodb: mongodb.New("27017", "test", "test"),
	}
}
