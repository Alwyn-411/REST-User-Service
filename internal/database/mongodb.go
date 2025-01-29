package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	DB *mongo.Database
}

func NewMongoDBConnection(uri string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("users")

	return &MongoDB{
		Client: client,
		DB:     db,
	}, nil
}

func (m *MongoDB) InsertOne(ctx context.Context,collectionName string, document interface{})(
	primitive.ObjectID,
	error,
) {
	collection := m.DB.Collection(collectionName)

	// Insert the document into the collection
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to insert document: %w", err)
	}

	// Get the inserted ID
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to get inserted ID")
	}

	return id, nil
}

func (m *MongoDB) FindOne(ctx context.Context, collectionName string, filter bson.M, result interface{}) error {
    collection := m.DB.Collection(collectionName)
    err := collection.FindOne(ctx, filter).Decode(result)
    if err == mongo.ErrNoDocuments {
        return fmt.Errorf("no document found with given criteria")
    }
    if err != nil {
        return fmt.Errorf("error in FindOne operation: %w", err)
    }
    return nil
}

func (m *MongoDB) FindAll(ctx context.Context, collectionName string, filter bson.M) ([]bson.M, error) {
    collection := m.DB.Collection(collectionName)
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, fmt.Errorf("error in FindAll operation: %w", err)
    }
    defer cursor.Close(ctx)

    var results []bson.M
    if err = cursor.All(ctx, &results); err != nil {
        return nil, fmt.Errorf("error decoding FindAll results: %w", err)
    }
    
    return results, nil
}

func (m *MongoDB) UpdateOne(ctx context.Context, collectionName string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
    collection := m.DB.Collection(collectionName)
    result, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return nil, fmt.Errorf("error in UpdateOne operation: %w", err)
    }
    if result.MatchedCount == 0 {
        return nil, fmt.Errorf("no document found to update")
    }
    
    return result, nil
}

func (m *MongoDB) DeleteOne(ctx context.Context, collectionName string, filter bson.M) (*mongo.DeleteResult, error) {
    collection := m.DB.Collection(collectionName)
    result, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return nil, fmt.Errorf("error in DeleteOne operation: %w", err)
    }
    if result.DeletedCount == 0 {
        return nil, fmt.Errorf("no document found to delete")
    }
    
    return result, nil
}


func (m *MongoDB) Close(context context.Context) error {
	return m.Client.Disconnect(context)
}