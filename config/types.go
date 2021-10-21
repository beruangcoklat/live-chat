package config

type Config struct {
	Port              string   `json:"port"`
	CassandraUsername string   `json:"cassandra_username"`
	CassandraPassword string   `json:"cassandra_password"`
	CassandraKeyspace string   `json:"cassandra_keyspace"`
	CassandraCluster  string   `json:"cassandra_cluster"`
	KafkaBroker       []string `json:"kafka_broker"`
}
