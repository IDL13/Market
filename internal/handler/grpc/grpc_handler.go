package grpc

import (
	"context"

	"github.com/IDL13/Market/internal/entity"
	"github.com/IDL13/Market/internal/usecase"
	api "github.com/IDL13/Market/pkg/api/proto"
)

func New() GRPCServer {
	return GRPCServer{
		usecase: usecase.NewGRPCUseCase(),
	}
}

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedMarketServer
	usecase usecase.GRPCUseCase
}

// ResetStatistics ...
func (s *GRPCServer) ResetStatistics(ctx context.Context, in *api.ResetStatisticsRequest) (*api.ResetStatisticsResponse, error) {
	status := s.usecase.Reset()

	return &api.ResetStatisticsResponse{Status: int32(status)}, nil
}

// SaveStatistics ...
func (s *GRPCServer) SaveStatistics(ctx context.Context, in *api.SaveStatisticsRequest) (*api.SaveStatisticsResponse, error) {
	status := s.usecase.GRPCSave(in.Date, in.Views, in.Clicks, in.Cost)

	return &api.SaveStatisticsResponse{Status: status}, nil
}

// ShowStatistics ...
func (s *GRPCServer) ShowStatistics(ctx context.Context, in *api.ShowStatisticsRequest) (*api.ShowStatisticsResponse, error) {
	statistics := s.usecase.GRPCShow(in.From, in.To)

	grpcStatistics := entity.StatisticsConvertor(statistics)

	return &api.ShowStatisticsResponse{Statistics: grpcStatistics}, nil
}

// mustEmbedUnimplementedMarketServer ...
func (s *GRPCServer) mustEmbedUnimplementedMarketServer() {}
