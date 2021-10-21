package domain

import (
	"context"
)

type (
	ChannelUsecase interface {
		Get(ctx context.Context) ([]Channel, error)
	}

	ChannelRepository interface {
		Get(ctx context.Context) ([]Channel, error)
	}
)

type (
	Channel struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
