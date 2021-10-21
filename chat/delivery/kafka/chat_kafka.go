package http

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/beruangcoklat/live-chat/config"
	"github.com/beruangcoklat/live-chat/constant"
	"github.com/beruangcoklat/live-chat/domain"
	"github.com/segmentio/kafka-go"
)

type chatHandler struct {
	chatUc domain.ChatUsecase
}

func New(chatUc domain.ChatUsecase) {
	handler := &chatHandler{
		chatUc: chatUc,
	}

	ctx := context.Background()
	go handler.consumeChat(ctx)
}

func (h *chatHandler) consumeChat(ctx context.Context) {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: config.GetConfig().KafkaBroker,
		Topic:   constant.KafkaTopicChat,
	})

	kafkaReader.SetOffsetAt(ctx, time.Now())

	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		var chat domain.Chat
		err = json.Unmarshal(msg.Value, &chat)
		if err != nil {
			log.Printf("[consumeChat] err=%v", err.Error())
			continue
		}

		h.chatUc.Broadcast(ctx, chat)
	}
}
