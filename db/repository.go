package db

import (
	"context"
	"golinebot/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) AddMessage(msg *model.Receive) error {
	bson, _ := toDoc(msg)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repo.db.Database("linedb").Collection("messages").InsertOne(ctx, bson)
	if err != nil {
		return err
	}
	return nil
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}
