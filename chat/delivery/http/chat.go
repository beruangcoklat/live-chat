package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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

	router.HandleFunc("/chat/{group_id}", handler.view).Methods(http.MethodGet)
	router.HandleFunc("/chat", handler.post).Methods(http.MethodPost)
}

func (h *chatHandler) view(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupIDStr, ok := vars["group_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	h.chatUc.NewClient(ctx, w, groupID)
	<-r.Context().Done()
	h.chatUc.CloseClient(ctx, w)
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
