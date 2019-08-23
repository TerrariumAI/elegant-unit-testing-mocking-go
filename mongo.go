package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// DataAccessLayer is an interface that models how we interract our database
type DataAccessLayer interface {
	CreateUser(user User) (string, error)
	GetUser(id string) (User, error)
	// DeleteUser(id int) (int, error)
}

// MongoDAL implements Mongo data access
type MongoDAL struct {
	client *mongo.Client
	dbName string
}

// NewMongoDAL initializes a MongoDAL and makes all necessary connections
func NewMongoDAL(uri string, dbName string) (MongoDAL, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return MongoDAL{}, err
	}

	return MongoDAL{
		client: client,
		dbName: dbName,
	}, nil
}

// CreateUser takes in a user and inserts it into the database
func (m *MongoDAL) CreateUser(user User) (string, error) {
	collection := m.client.Database(m.dbName).Collection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

// GetUser takes an id and returns the resulting user
func (m *MongoDAL) GetUser(id string) (User, error) {
	collection := m.client.Database(m.dbName).Collection("users")
	objID, _ := primitive.ObjectIDFromHex(id)
	println(objID.String())
	var user User
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	user.ID = id
	return user, nil
}
