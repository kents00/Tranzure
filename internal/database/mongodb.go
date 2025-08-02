package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/kento/tranzure/internal/config"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDB(cfg *config.MongoDBConfig) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create client options
	clientOpts := options.Client().ApplyURI(cfg.GetConnectionString())

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	database := client.Database(cfg.Database)

	return &MongoDB{
		Client:   client,
		Database: database,
	}, nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

func (m *MongoDB) Health(ctx context.Context) error {
	return m.Client.Ping(ctx, readpref.Primary())
}

// Collections returns commonly used collections
func (m *MongoDB) Collections() *Collections {
	return &Collections{
		Users:        m.Database.Collection("users"),
		Wallets:      m.Database.Collection("wallets"),
		Transactions: m.Database.Collection("transactions"),
		KYC:          m.Database.Collection("kyc_verifications"),
		Sessions:     m.Database.Collection("sessions"),
		AuditLogs:    m.Database.Collection("audit_logs"),
	}
}

type Collections struct {
	Users        *mongo.Collection
	Wallets      *mongo.Collection
	Transactions *mongo.Collection
	KYC          *mongo.Collection
	Sessions     *mongo.Collection
	AuditLogs    *mongo.Collection
}
