package domain

import (
	"context"
	"net/http"
)

type (
	ChatUsecase interface {
		NewClient(ctx context.Context, r http.ResponseWriter, groupID int64)
		CloseClient(ctx context.Context, r http.ResponseWriter)
		Post(ctx context.Context, chat Chat) error
	}
)

type (
	Chat struct {
		GroupID int64  `json:"group_id"`
		Message string `json:"message"`
	}
)
