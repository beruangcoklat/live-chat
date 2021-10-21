package cassandra

import (
	"context"

	"github.com/beruangcoklat/live-chat/domain"
	"github.com/gocql/gocql"
)

type repository struct {
	cassandraSession *gocql.Session
}

func New(cassandraSession *gocql.Session) domain.ChannelRepository {
	return &repository{
		cassandraSession: cassandraSession,
	}
}

func (r *repository) Get(ctx context.Context) ([]domain.Channel, error) {
	query := `SELECT
		id,
		name
	FROM
		channel`

	channel := domain.Channel{}
	result := []domain.Channel{}
	iter := r.cassandraSession.Query(query).Iter()
	for iter.Scan(&channel.ID, &channel.Name) {
		result = append(result, channel)
	}
	return result, nil
}
