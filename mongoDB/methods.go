package mongoDB

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find executes a find command and returns a Cursor over the matching documents in the collection.
func (m *MongoImpl) Find(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOptions,
) (*mongo.Cursor, error) {
	metricName := "Find"
	startTime := time.Now()

	res, err := m.mongoCollection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, nil
}

// FindOne returns up to one document that matches the model.
func (m *MongoImpl) FindOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneOptions,
) *mongo.SingleResult {
	metricName := "FindOne"
	startTime := time.Now()

	res := m.mongoCollection.FindOne(ctx, filter, opts...)

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res
}

func (m *MongoImpl) FindOneAndUpdate(ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.FindOneAndUpdateOptions,
) *mongo.SingleResult {
	metricName := "FindOne"
	startTime := time.Now()

	res := m.mongoCollection.FindOneAndUpdate(ctx, filter, update, opts...)

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res
}

// InsertOne inserts a single document into the collection.
func (m *MongoImpl) InsertOne(
	ctx context.Context,
	document interface{},
	opts ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	metricName := "InsertOne"
	startTime := time.Now()

	res, err := m.mongoCollection.InsertOne(ctx, document, opts...)
	if err != nil {
		m.logger.Println("collection.InsertOne", "err", err)
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, err
}

func (m *MongoImpl) ReplaceOne(
	ctx context.Context,
	filter interface{},
	replacement interface{},
	opts ...*options.ReplaceOptions,
) (*mongo.UpdateResult, error) {
	metricName := "InsertOne"
	startTime := time.Now()

	res, err := m.mongoCollection.ReplaceOne(ctx, filter, replacement, opts...)
	if err != nil {
		m.logger.Println("collection.ReplaceOne", "err", err)
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, err
}

// InsertMany inserts a single document into the collection.
func (m *MongoImpl) InsertMany(
	ctx context.Context,
	documents []interface{},
	opts ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	metricName := "InsertMany"
	startTime := time.Now()
	res, err := m.mongoCollection.InsertMany(ctx, documents, opts...)
	if err != nil {
		m.logger.Println("collection.InsertMany", "err", err)
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, err
}

// UpdateOne updates a document
func (m *MongoImpl) UpdateOne(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	metricName := "UpdateOne"
	startTime := time.Now()

	res, err := m.mongoCollection.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		m.logger.Println("collection.UpdateOne", "err", err)
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, err
}

// DeleteOne deletes a document
func (m *MongoImpl) DeleteOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	metricName := "DeleteOne"
	startTime := time.Now()

	res, err := m.mongoCollection.DeleteOne(ctx, filter, opts...)
	if err != nil {
		m.logger.Println("collection.DeleteOne", "err", err)
		return nil, err
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, err
}

func (m *MongoImpl) DeleteMany(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	metricName := "DeleteMany"
	startTime := time.Now()

	res, err := m.mongoCollection.DeleteMany(ctx, filter, opts...)
	if err != nil {
		m.logger.Println("collection.DeleteMany", "err", err)
		return nil, err
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, err
}

func (m *MongoImpl) DropAll(ctx context.Context, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	metricName := "DropAll"
	startTime := time.Now()
	res, err := m.mongoCollection.Indexes().DropAll(ctx, opts...)
	if err != nil {
		m.logger.Println("collection.DropAll", "err", err)
		return nil, err
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, err

}

// UpdateMany updates all documents that matches the filter
func (m *MongoImpl) UpdateMany(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	metricName := "UpdateMany"
	startTime := time.Now()

	res, err := m.mongoCollection.UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		m.logger.Println("collection.UpdateMany", "err", err)
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "res", res, "elapsedMs", elapsed.Milliseconds())

	return res, err
}

// CountDocuments returns the number of documents that matches the filter
func (m *MongoImpl) CountDocuments(
	ctx context.Context,
	filter interface{},
	opts ...*options.CountOptions,
) (int64, error) {
	metricName := "CountDocuments"
	startTime := time.Now()

	count, err := m.mongoCollection.CountDocuments(ctx, filter, opts...)
	if err != nil {
		m.logger.Println("collection.CountDocuments", "err", err)
	}

	elapsed := time.Since(startTime)
	m.logger.Println(metricName, "count", count, "elapsedMs", elapsed.Milliseconds())

	return count, err
}
