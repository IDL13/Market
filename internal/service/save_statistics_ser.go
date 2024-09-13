package service

import "github.com/IDL13/Market/internal/entity"

func (s service) SaveStatistics(json entity.SaveRequest) int8 {
	ch := make(chan int8)

	go func(ch chan int8) {
		ch <- s.mongodb.AddStatistics(json)
	}(ch)

	if <-ch == 1 {
		return 1
	} else {
		return 0
	}
}
