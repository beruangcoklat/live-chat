package usecase

import (
	"context"

	"github.com/beruangcoklat/live-chat/domain"
)

type usecase struct {
	channelRepo domain.ChannelRepository
}

func New(channelRepo domain.ChannelRepository) domain.ChannelUsecase {
	return &usecase{
		channelRepo: channelRepo,
	}
}

func (uc *usecase) Get(ctx context.Context) ([]domain.Channel, error) {
	result, err := uc.channelRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
