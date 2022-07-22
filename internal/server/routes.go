package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/romik1505/communicationGraph/docs"
	"github.com/romik1505/communicationGraph/internal/app/mapper"
	"github.com/romik1505/communicationGraph/internal/app/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	cs *service.CommunicationService
}

func NewHandler(cs *service.CommunicationService) *Handler {
	return &Handler{
		cs: cs,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/record", h.addRecord).Methods("POST")
	router.HandleFunc("/graph", h.getGraph)
	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	return router
}

// @Summary Record
// @Tags record
// @Description add new communication node
// @ID record
// @Accept json
// @Produce json
// @Param input body mapper.JournalRecord true "account info"
// @Success 200
// @Failure 400,500 {string} string
// @Router /record [post]
func (h Handler) addRecord(w http.ResponseWriter, r *http.Request) {
	rec := mapper.JournalRecord{}
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.cs.AddRecord(r.Context(), rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Graph
// @Tags graph
// @Description return graph nodes links
// @ID graph
// @Accept json
// @Produce json
// @Success 200 {object} mapper.JournalRecord "status"
// @Failure 400,500 {string} string
// @Router /graph [get]
func (h Handler) getGraph(w http.ResponseWriter, r *http.Request) {
	res, err := h.cs.GetGraph(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
