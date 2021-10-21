package mongostorage

import (
	"context"
	"integrations/cmd/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoStorage containing connection
type MongoStorage struct {
	Config *config.DatabaseConfig
	Client *mongo.Client
}

// MongoDB database used by app
type MongoDB struct {
	Storage  *MongoStorage
	Database *mongo.Database
}

// NewMongoStorage returns new MongoStorage instance
func NewMongoStorage(c *config.DatabaseConfig) *MongoStorage {
	client, err := mongo.NewClient(options.Client().ApplyURI(c.ConnectionURL()))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		panic(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	mongoStorage := &MongoStorage{
		Config: c,
		Client: client,
	}
	return mongoStorage
}

// NewMongoDB returns new MongoDB struct instance
func NewMongoDB(c *config.DatabaseConfig) *MongoDB {
	storage := NewMongoStorage(c)
	db := storage.Client.Database(c.DBName)
	mongoDB := &MongoDB{
		Storage:  storage,
		Database: db,
	}
	return mongoDB
}

// NewSession creates new mongodb session
func (ms *MongoStorage) NewSession() (mongo.Session, error) {
	session, err := ms.Client.StartSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

// NewTransaction starts a new transaction
func (ms *MongoStorage) NewTransaction(s mongo.Session) error {
	err := s.StartTransaction()
	return err
}
