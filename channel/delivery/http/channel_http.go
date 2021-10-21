package http

import (
	"encoding/json"
	"net/http"

	"github.com/beruangcoklat/live-chat/domain"
	"github.com/gorilla/mux"
)

type channelHandler struct {
	channelUc domain.ChannelUsecase
}

func New(router *mux.Router, channelUc domain.ChannelUsecase) {
	handler := &channelHandler{
		channelUc: channelUc,
	}

	router.HandleFunc("/channel", handler.get).Methods(http.MethodGet)
}

func (h *channelHandler) get(w http.ResponseWriter, r *http.Request) {
	result, err := h.channelUc.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
