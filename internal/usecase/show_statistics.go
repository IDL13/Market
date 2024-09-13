package usecase

import (
	"encoding/json"
	"io"

	"github.com/IDL13/Market/internal/entity"
	"github.com/IDL13/Market/internal/service"
)

type showStatistics struct{}

func (s showStatistics) Show(request io.ReadCloser) []entity.Statistics {
	var jsonBody entity.ShowRequestAdapter

	service := service.NewService()
	ch := make(chan []entity.Statistics)

	json.NewDecoder(request).Decode(&jsonBody)

	go func(ch chan []entity.Statistics) {
		ch <- service.ShowStatistics(entity.NewShowRequest(jsonBody))
	}(ch)

	return <-ch
}

func (s showStatistics) GRPCShow(from, to string) []entity.Statistics {
	adapter := entity.ShowRequestAdapter{
		From: from,
		To:   to,
	}

	showRequest := entity.NewShowRequest(adapter)

	service := service.NewService()
	ch := make(chan []entity.Statistics)

	go func(ch chan []entity.Statistics) {
		ch <- service.ShowStatistics(showRequest)
	}(ch)

	return <-ch
}
