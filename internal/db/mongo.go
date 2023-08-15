package db

import (
	"context"
	"errors"
	"go-clean-architecture/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slog"
	"time"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

func MustConnect(opts config.DatabaseOpts) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(opts.Uri))
	if err != nil {
		return err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return errors.New("cannot connect to database host")
	}

	database = client.Database(opts.Name)

	slog.Debug("Database connected", "name", opts.Name)
	return nil
}

func Close() {
	if client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = client.Disconnect(ctx)
}

func Collection(name string) *mongo.Collection {
	return database.Collection(name)
}
