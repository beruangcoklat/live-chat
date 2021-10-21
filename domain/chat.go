package domain

import (
	"context"
	"net/http"
)

type (
	ChatUsecase interface {
		NewClient(ctx context.Context, r http.ResponseWriter, groupID string)
		CloseClient(ctx context.Context, r http.ResponseWriter, groupID string)
		Post(ctx context.Context, chat Chat) error
		Get(ctx context.Context, channelID string, createdAt int64, limit int) ([]Chat, error)
		Broadcast(ctx context.Context, chat Chat)
	}

	ChatRepository interface {
		Get(ctx context.Context, channelID string, createdAt int64, limit int) ([]Chat, error)
		Create(ctx context.Context, chat Chat) error
		Publish(ctx context.Context, chat Chat) error
	}
)

type (
	Chat struct {
		ChannelID string `json:"channel_id"`
		Sender    string `json:"sender"`
		Message   string `json:"message"`
		CreatedAt int64  `json:"created_at"`
	}
)
