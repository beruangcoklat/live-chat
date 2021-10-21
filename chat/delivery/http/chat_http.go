package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/beruangcoklat/live-chat/domain"
	"github.com/gorilla/mux"
)

type chatHandler struct {
	chatUc domain.ChatUsecase
}

func New(router *mux.Router, chatUc domain.ChatUsecase) {
	handler := &chatHandler{
		chatUc: chatUc,
	}

	router.HandleFunc("/chat/{channel_id}", handler.view).Methods(http.MethodGet)
	router.HandleFunc("/chat", handler.post).Methods(http.MethodPost)
	router.HandleFunc("/chat", handler.get).Methods(http.MethodGet)
}

func (h *chatHandler) view(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID, ok := vars["channel_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	h.chatUc.NewClient(ctx, w, channelID)
	<-r.Context().Done()
	h.chatUc.CloseClient(ctx, w, channelID)
}

func (h *chatHandler) post(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var chat domain.Chat
	err = json.Unmarshal(bytes, &chat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.chatUc.Post(r.Context(), chat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *chatHandler) get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	channelID := query.Get("channel_id")

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdAt, err := strconv.ParseInt(query.Get("created_at"), 10, 64)
	if err != nil {
		createdAt = time.Now().UnixMilli()
	}

	result, err := h.chatUc.Get(r.Context(), channelID, createdAt, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
