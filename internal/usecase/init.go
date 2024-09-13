package usecase

import (
	"io"

	"github.com/IDL13/Market/internal/entity"
)

//go:generate mockgen -source=init.go -destination=../mock/mock.go -package=mock
type UseCase interface {
	reset
	save
	show
}

//go:generate mockgen -source=init.go -destination=../mock_grpc/grpc_mock.go -package=mock_grpc
type GRPCUseCase interface {
	Reset() int8
	GRPCSave(date string, views int32, clicks int32, coast float32) int32
	GRPCShow(from, to string) []entity.Statistics
}

type reset interface {
	Reset() int8
}

type save interface {
	Save(request io.ReadCloser) int8
}

type show interface {
	Show(request io.ReadCloser) []entity.Statistics
}

type useCase struct {
	resetStatistics
	saveStatistics
	showStatistics
}

func NewUseCase() UseCase {
	return useCase{}
}

func NewGRPCUseCase() GRPCUseCase {
	return useCase{}
}
