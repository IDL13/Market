package service

import (
	"github.com/IDL13/Market/internal/entity"
)

func (s service) ShowStatistics(json entity.ShowRequest) []entity.Statistics {
	ch := make(chan []entity.Statistics)

	go func(ch chan []entity.Statistics) {
		ch <- s.mongodb.GetStatistics(json)
	}(ch)

	return <-ch
}
