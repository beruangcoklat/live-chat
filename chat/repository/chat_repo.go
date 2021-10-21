package cassandra

import (
	"context"
	"encoding/json"

	"github.com/beruangcoklat/live-chat/constant"
	"github.com/beruangcoklat/live-chat/domain"
	"github.com/gocql/gocql"
	"github.com/segmentio/kafka-go"
)

type repository struct {
	cassandraSession *gocql.Session
	kafkaWriter      *kafka.Writer
}

func New(cassandraSession *gocql.Session, kafkaWriter *kafka.Writer) domain.ChatRepository {
	return &repository{
		cassandraSession: cassandraSession,
		kafkaWriter:      kafkaWriter,
	}
}

func (r *repository) Get(ctx context.Context, channelID string, createdAt int64, limit int) ([]domain.Chat, error) {
	query := `SELECT
		channel_id,
		sender,
		message,
		created_at
	FROM
		chat
	WHERE
		channel_id = ?
		and created_at < ?
	ORDER BY
		created_at DESC
	LIMIT ?`

	chat := domain.Chat{}
	result := []domain.Chat{}
	iter := r.cassandraSession.Query(query, channelID, createdAt, limit).Iter()
	for iter.Scan(&chat.ChannelID, &chat.Sender, &chat.Message, &chat.CreatedAt) {
		result = append(result, chat)
	}
	return result, nil
}

func (r *repository) Create(ctx context.Context, chat domain.Chat) error {
	query := `INSERT INTO chat (channel_id, sender, message, created_at) VALUES (?, ?, ?, ?)`
	err := r.cassandraSession.Query(query, chat.ChannelID, chat.Sender, chat.Message, chat.CreatedAt).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Publish(ctx context.Context, chat domain.Chat) error {
	value, err := json.Marshal(chat)
	if err != nil {
		return err
	}

	err = r.kafkaWriter.WriteMessages(ctx, kafka.Message{
		Topic: constant.KafkaTopicChat,
		Value: value,
	})
	if err != nil {
		return err
	}

	return nil
}
