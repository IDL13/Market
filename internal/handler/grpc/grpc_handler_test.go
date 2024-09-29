package grpc

import (
	"testing"

	"github.com/IDL13/Market/internal/entity"
	"github.com/IDL13/Market/internal/mock_grpc"
	api "github.com/IDL13/Market/pkg/api/proto"
	"github.com/golang/mock/gomock"
)

func TestResetStatistics(t *testing.T) {
	tests := []struct {
		name string
		f    func(s *mock_grpc.MockGRPCUseCase)
	}{{
		name: "OK",
		f: func(s *mock_grpc.MockGRPCUseCase) {
			s.EXPECT().Reset().Return(int8(1)).AnyTimes()
		},
	},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_grpc.NewMockGRPCUseCase(ctrl)
			test.f(srvc)
		})
	}
}

func TestSaveStatistics(t *testing.T) {
	tests := []struct {
		name      string
		inputBody api.SaveStatisticsRequest
		f         func(s *mock_grpc.MockGRPCUseCase, request api.SaveStatisticsRequest)
	}{{
		name: "OK",
		inputBody: api.SaveStatisticsRequest{
			Date:   "28.09.2024",
			Views:  10,
			Clicks: 10,
			Cost:   10,
		},
		f: func(s *mock_grpc.MockGRPCUseCase, request api.SaveStatisticsRequest) {
			s.EXPECT().GRPCSave(request.Date, request.Views, request.Clicks, request.Cost).Return(int32(1)).AnyTimes()
		},
	},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_grpc.NewMockGRPCUseCase(ctrl)
			test.f(srvc, test.inputBody)
		})
	}
}

func TestShowStatistics(t *testing.T) {
	tests := []struct {
		name      string
		inputBody api.ShowStatisticsRequest
		f         func(s *mock_grpc.MockGRPCUseCase, request api.ShowStatisticsRequest)
	}{{
		name: "OK",
		inputBody: api.ShowStatisticsRequest{
			From: "28-09-2024",
			To:   "28-09-2024",
		},
		f: func(s *mock_grpc.MockGRPCUseCase, request api.ShowStatisticsRequest) {
			s.EXPECT().GRPCShow(request.From, request.To).Return([]entity.Statistics{}).AnyTimes()
		},
	},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_grpc.NewMockGRPCUseCase(ctrl)
			test.f(srvc, test.inputBody)
		})
	}
}
