package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/beruangcoklat/live-chat/domain"
)

type usecase struct {
	notifier       chan *domain.Chat
	newClients     chan http.ResponseWriter
	closingClients chan http.ResponseWriter
	clients        map[http.ResponseWriter]struct{}
}

func New() domain.ChatUsecase {
	uc := &usecase{
		notifier:       make(chan *domain.Chat),
		newClients:     make(chan http.ResponseWriter),
		closingClients: make(chan http.ResponseWriter),
		clients:        make(map[http.ResponseWriter]struct{}),
	}

	go func() {
		for {
			select {
			case s := <-uc.newClients:
				uc.clients[s] = struct{}{}
			case s := <-uc.closingClients:
				delete(uc.clients, s)
			case event := <-uc.notifier:
				for w := range uc.clients {
					uc.sendData(w, event)
				}
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			uc.notifier <- &domain.Chat{
				GroupID: 1,
				Message: fmt.Sprintf("%v", time.Now().Unix()),
			}
		}
	}()

	return uc
}

func (uc *usecase) NewClient(ctx context.Context, w http.ResponseWriter, groupID int64) {
	uc.newClients <- w
}

func (uc *usecase) CloseClient(ctx context.Context, w http.ResponseWriter) {
	uc.closingClients <- w
}

func (uc *usecase) sendData(w http.ResponseWriter, event *domain.Chat) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bytes, err := json.Marshal(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "data: %s\n\n", string(bytes))
	flusher.Flush()
}

func (uc *usecase) Post(ctx context.Context, chat domain.Chat) error {
	uc.notifier <- &chat
	return nil
}
