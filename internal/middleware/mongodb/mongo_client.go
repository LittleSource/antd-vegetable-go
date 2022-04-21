package mongodb

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db *mongo.Database
)

func InitMongodb() {
	ctx := context.Background()
	uri, err := gcfg.Instance().Get(ctx, "database.mongodb.uri")
	if err != nil {
		g.Log().Errorf(ctx, "Mongo uri error: ", err)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri.String()))
	if err != nil {
		g.Log().Errorf(ctx, "Mongo connect error: ", err)
	} else {
		g.Log().Info(ctx, "Mongodb connect succeeded!")
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		g.Log().Errorf(ctx, "Mongo ping error: ", err)
	} else {
		g.Log().Notice(ctx, "Mongodb ping succeeded!")
	}
	database, err := gcfg.Instance().Get(ctx, "database.mongodb.database")
	if err != nil {
		g.Log().Errorf(ctx, "Mongo uri error: ", err)
	}
	db = client.Database(database.String())
}

func Database() *mongo.Database {
	return db
}
