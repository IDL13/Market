package http

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/IDL13/Market/internal/entity"
	mock_serv "github.com/IDL13/Market/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSaveStatistics(t *testing.T) {
	tests := []struct {
		name         string
		inputService string
		inputBody    string
		expectedCode int
		f            func(s *mock_serv.MockUseCase, u io.ReadCloser)
		expectBody   string
	}{
		{
			name:         "Ok",
			inputService: "test",
			inputBody: `{
				"date":"2020-12-12",
				"views": 10,
				"clicks": 10,
				"cost": 10.0	
			}`,
			expectedCode: 200,
			f: func(s *mock_serv.MockUseCase, u io.ReadCloser) {
				s.EXPECT().Save(u).Return(int8(1)).AnyTimes()
			},
			expectBody: `{"status_code":200,"message":"Request completed successfully"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockUseCase(ctrl)
			readerCloser := io.NopCloser(strings.NewReader(test.inputBody))
			test.f(srvc, readerCloser)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/save", service.SaveStatistics)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/save", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestResetStatistics(t *testing.T) {
	tests := []struct {
		name         string
		expectedCode int
		f            func(s *mock_serv.MockUseCase)
		expectBody   string
	}{
		{
			name:         "Ok",
			expectedCode: 200,
			f: func(s *mock_serv.MockUseCase) {
				s.EXPECT().Reset().Return(int8(1)).AnyTimes()
			},
			expectBody: `{"status_code":200,"message":"Request completed successfully"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockUseCase(ctrl)
			test.f(srvc)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/reset", service.ResetStatistics)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/reset", bytes.NewBufferString(""))

			r.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestShowStatistics(t *testing.T) {
	tests := []struct {
		name         string
		inputService string
		inputBody    string
		expectedCode int
		f            func(s *mock_serv.MockUseCase, u io.ReadCloser)
		expectBody   string
	}{
		{
			name:         "Ok",
			inputService: "test",
			inputBody:    `{"from":"2020-12-12","to": "2020-12-12"}`,
			expectedCode: 200,
			f: func(s *mock_serv.MockUseCase, u io.ReadCloser) {
				s.EXPECT().Show(u).Return([]entity.Statistics{}).AnyTimes()
			},
			expectBody: `{"status_code":200,"message":"{\"statistics\":null}"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_serv.NewMockUseCase(ctrl)
			readerCloser := io.NopCloser(strings.NewReader(test.inputBody))
			test.f(srvc, readerCloser)

			service := New()

			r := http.NewServeMux()
			r.HandleFunc("/show", service.ShowStatistics)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/show", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)
			log.Println(w.Code)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
