package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

var configuration = New()

func NewMongoDatabase() *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, _ := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	mongoPoolMax, _ := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	mongoMaxIdleTime, _ := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))

	option := options.Client().
		ApplyURI(ClientUrl()).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		panic(err)
	}

	if err := client.Connect(ctx); err != nil {
		panic(err)
	}
	database := client.Database(configuration.Get("MONGO_DATABASE"))
	fmt.Println(client.Ping(ctx, nil))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func ClientUrl() string {
	connPattern := "mongodb://%v:%v@%v:%v"
	if configuration.Get("MONGO_USERNAME") == "" {
		connPattern = "mongodb://%s%s%v:%v"
	}

	clientUrl := fmt.Sprintf(connPattern,
		configuration.Get("MONGO_USERNAME"),
		configuration.Get("MONGO_PASSWORD"),
		configuration.Get("MONGO_HOST"),
		configuration.Get("MONGO_PORT"),
	)
	return clientUrl
}
