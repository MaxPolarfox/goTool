package mongoDB

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo interface {
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, documents []interface{},
		opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{},
		opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	UpdateMany(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...*options.UpdateOptions,
	) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error)
	CountDocuments(
		ctx context.Context,
		filter interface{},
		opts ...*options.CountOptions,
	) (int64, error)
}

type MongoImpl struct {
	logger          log.Logger
	mongoCollection *mongo.Collection
}

func NewMongo(opts Options, logger log.Logger) Mongo {
	logger.Println("Instantiating MongoDB", "options", opts)

	ctx, _ := context.WithTimeout(context.Background(), 25*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(opts.Connection))
	if err != nil {
		logger.Fatalf("NewMongo.error", "err", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Fatalf("Mongo Ping failed", "err", err)
	}
	collection := client.Database(opts.Name).Collection(opts.Collection)

	//always create id index
	idIndex := mongo.IndexModel{
		Keys: bson.M{
			"id": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(ctx, idIndex)
	if err != nil {
		logger.Println("CreatingIndex", "err", err)
	}

	return &MongoImpl{
		mongoCollection: collection,
		logger:          logger,
	}
}
