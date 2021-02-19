package mongoDB

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMongo(t *testing.T) {

	testIntegrationOptions := Options{
		Name:       "mongodb-test",
		Collection: "mongodb-test-collection",
		Connection: "mongodb://localhost:27017/goTools-development?retryWrites=true&w=majority",
	}

	testClient := NewMongo(testIntegrationOptions)

	Convey("InsertOne, FindOne, DeleteOne", t, func() {
		randomName := "test-mongo-collection"

		body := bson.M{"name": randomName, "age": 5}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		res, err := testClient.InsertOne(ctx, body)
		if err != nil {
			t.Error(err)
		}

		So(res.InsertedID, ShouldNotBeNil)

		var result map[string]interface{}
		getError := testClient.FindOne(ctx, bson.M{"name": randomName}).Decode(&result)
		if getError != nil {
			t.Error("error ", getError)
		}
		So(result["name"], ShouldEqual, randomName)
		So(result["age"], ShouldEqual, 5)

		// DeleteOne test
		deleteResult, deleteError := testClient.DeleteOne(ctx, bson.M{"name": randomName})
		So(deleteError, ShouldBeNil)
		So(deleteResult.DeletedCount, ShouldEqual, 1)

	})
}
