package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

const uri = "mongodb://localhost:27017"

var Client = connection{}

type connection struct {
	client *mongo.Client
}

func (c *connection) CreateConnection(ctx context.Context) {
	var err error
	c.client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}

}

func (c *connection) PingMongo(ctx context.Context) {

	err := c.client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Pinged mongo!")

}

func (c *connection) Disconnect(ctx context.Context) {
	err := c.client.Disconnect(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Disconnected!")
}
