package http

import (
	"encoding/json"
	"net/http"

	"github.com/IDL13/Market/internal/entity"
	"github.com/IDL13/Market/internal/usecase"
	resp "github.com/IDL13/Market/pkg/response"
)

type Handler interface {
	SaveStatistics(w http.ResponseWriter, r *http.Request)
	ShowStatistics(w http.ResponseWriter, r *http.Request)
	ResetStatistics(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	usecase usecase.UseCase
}

func New() Handler {
	return &handler{
		usecase: usecase.NewUseCase(),
	}
}

func (h handler) SaveStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "PATCH" || r.Method == "PUT" || r.Method == "DELETE" {
		w.WriteHeader(resp.Status404)
		w.Write(resp.New(resp.Status404, resp.Message400))
	} else {
		if usecase := h.usecase.Save(r.Body); usecase != 1 {
			w.WriteHeader(resp.Status500)
			w.Write(resp.New(resp.Status500, resp.Message500))
		} else {
			w.WriteHeader(resp.Status200)
			w.Write(resp.New(resp.Status200, resp.Message200))
		}
	}
}

func (h handler) ShowStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "PATCH" || r.Method == "PUT" || r.Method == "DELETE" {
		w.WriteHeader(resp.Status404)
		w.Write(resp.New(resp.Status400, resp.Message400))
	} else {
		usecase := entity.StatisticsResponse{
			Statistics: h.usecase.Show(r.Body),
		}

		respBody, err := json.Marshal(usecase)

		if err != nil {
			w.WriteHeader(resp.Status500)
			w.Write(resp.New(resp.Status500, resp.Message500))
		}

		w.WriteHeader(resp.Status200)
		w.Write(resp.New(resp.Status200, string(respBody)))
	}
}

func (h handler) ResetStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "PATCH" || r.Method == "PUT" || r.Method == "DELETE" {
		w.WriteHeader(resp.Status404)
		w.Write(resp.New(resp.Status404, resp.Message400))
	} else {
		if usecase := h.usecase.Reset(); usecase != 1 {
			w.WriteHeader(resp.Status500)
			w.Write(resp.New(resp.Status500, resp.Message500))
		} else {
			w.WriteHeader(resp.Status200)
			w.Write(resp.New(resp.Status200, resp.Message200))
		}
	}
}
