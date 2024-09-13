package usecase

import (
	"encoding/json"
	"io"

	"github.com/IDL13/Market/internal/entity"
	"github.com/IDL13/Market/internal/service"
)

type saveStatistics struct{}

func (s saveStatistics) Save(request io.ReadCloser) int8 {
	var jsonBody entity.SaveReuqestAdapter

	service := service.NewService()
	ch := make(chan int8)

	json.NewDecoder(request).Decode(&jsonBody)

	go func(ch chan int8) {
		ch <- service.SaveStatistics(entity.NewSaveRequest(jsonBody))
	}(ch)

	if <-ch == 1 {
		return 1
	} else {
		return 0
	}
}

func (s saveStatistics) GRPCSave(date string, views int32, clicks int32, coast float32) int32 {
	adapter := entity.SaveReuqestAdapter{
		Date:   date,
		Views:  views,
		Clicks: clicks,
		Cost:   coast,
	}

	saveRequest := entity.NewSaveRequest(adapter)

	service := service.NewService()
	ch := make(chan int32)

	go func(ch chan int32) {
		ch <- int32(service.SaveStatistics(saveRequest))
	}(ch)

	if <-ch == 1 {
		return 1
	} else {
		return 0
	}
}
