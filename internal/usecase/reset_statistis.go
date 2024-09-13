package usecase

import "github.com/IDL13/Market/internal/service"

type resetStatistics struct{}

func (r resetStatistics) Reset() int8 {
	service := service.NewService()
	ch := make(chan int8)

	go func(ch chan int8) {
		ch <- service.ResetStatistics()
	}(ch)

	if <-ch == 1 {
		return 1
	} else {
		return 0
	}
}
