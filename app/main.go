package main

import (
	"log"
	"net/http"

	chathttp "github.com/beruangcoklat/live-chat/chat/delivery/http"
	chatkafka "github.com/beruangcoklat/live-chat/chat/delivery/kafka"
	chatrepository "github.com/beruangcoklat/live-chat/chat/repository"
	chatusecase "github.com/beruangcoklat/live-chat/chat/usecase"
	"github.com/beruangcoklat/live-chat/config"
	"github.com/beruangcoklat/live-chat/domain"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

var (
	chatRepo domain.ChatRepository
	chatUc   domain.ChatUsecase

	cassandraSession *gocql.Session
	kafkaWriter      *kafka.Writer
)

func initConfig() error {
	return config.Init("/etc/live-chat/config.json")
}

func initCassandra() (err error) {
	cfg := config.GetConfig()
	cluster := gocql.NewCluster(cfg.CassandraCluster)
	cluster.Keyspace = cfg.CassandraKeyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: cfg.CassandraUsername, Password: cfg.CassandraPassword}
	cassandraSession, err = cluster.CreateSession()
	if err != nil {
		return err
	}
	return nil
}

func initKafka() {
	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: config.GetConfig().KafkaBroker,
	})
}

func initRepo() {
	chatRepo = chatrepository.New(cassandraSession, kafkaWriter)
}

func initUsecase() {
	chatUc = chatusecase.New(chatRepo)
}

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = initCassandra()
	if err != nil {
		log.Fatal(err)
	}

	initKafka()

	defer func() {
		cassandraSession.Close()
		err = kafkaWriter.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	initRepo()
	initUsecase()

	router := mux.NewRouter()
	chathttp.New(router, chatUc)
	chatkafka.New(chatUc)

	port := config.GetConfig().Port
	log.Print("listen :" + port)
	http.ListenAndServe(":"+port, router)
}
