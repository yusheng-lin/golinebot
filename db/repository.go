package db

import (
	"context"
	"golinebot/model"
	"golinebot/service"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) service.IRepository {
	var repo service.IRepository = &Repository{
		db: db,
	}
	return repo
}

func (repo *Repository) AddMessage(msg *model.Message) error {
	bson, _ := toDoc(msg)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repo.db.Database("linedb").Collection("messages").InsertOne(ctx, bson)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetMessages(lineuserId string) (*[]model.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := repo.db.Database("linedb").Collection("messages").Find(ctx, bson.D{{Key: "lineuserid", Value: lineuserId}})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	messages := []model.Message{}

	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		doc, err := bson.Marshal(result)
		var m model.Message
		err = bson.Unmarshal(doc, &m)
		messages = append(messages, m)
	}

	return &messages, nil
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}
