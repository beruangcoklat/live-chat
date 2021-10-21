package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/beruangcoklat/live-chat/domain"
)

type usecase struct {
	channelMap sync.Map
	chatRepo   domain.ChatRepository
}

type channel struct {
	notifier       chan *domain.Chat
	newClients     chan http.ResponseWriter
	closingClients chan http.ResponseWriter
	clients        map[http.ResponseWriter]struct{}
}

func newChannel() *channel {
	return &channel{
		notifier:       make(chan *domain.Chat),
		newClients:     make(chan http.ResponseWriter),
		closingClients: make(chan http.ResponseWriter),
		clients:        make(map[http.ResponseWriter]struct{}),
	}
}

func New(chatRepo domain.ChatRepository) domain.ChatUsecase {
	uc := &usecase{
		chatRepo: chatRepo,
	}

	go func() {
		for {
			uc.channelMap.Range(func(key, value interface{}) bool {
				data, ok := value.(*channel)
				if !ok {
					return true
				}

				select {
				case event := <-data.notifier:
					for w := range data.clients {
						uc.sendData(w, event)
					}

				case s := <-data.newClients:
					data.clients[s] = struct{}{}

				case s := <-data.closingClients:
					delete(data.clients, s)
				}
				return true
			})
		}
	}()

	return uc
}

func (uc *usecase) NewClient(ctx context.Context, w http.ResponseWriter, channelID string) {
	dataItf, _ := uc.channelMap.LoadOrStore(channelID, newChannel())
	data := dataItf.(*channel)
	data.newClients <- w
}

func (uc *usecase) CloseClient(ctx context.Context, w http.ResponseWriter, channelID string) {
	dataItf, _ := uc.channelMap.Load(channelID)
	data := dataItf.(*channel)
	data.closingClients <- w
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
	chat.CreatedAt = time.Now().UnixMilli()

	err := uc.chatRepo.Create(ctx, chat)
	if err != nil {
		return err
	}

	err = uc.chatRepo.Publish(ctx, chat)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Get(ctx context.Context, channelID string, createdAt int64, limit int) ([]domain.Chat, error) {
	result, err := uc.chatRepo.Get(ctx, channelID, createdAt, limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *usecase) Broadcast(ctx context.Context, chat domain.Chat) {
	dataItf, _ := uc.channelMap.LoadOrStore(chat.ChannelID, newChannel())
	data := dataItf.(*channel)
	data.notifier <- &chat
}
