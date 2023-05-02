package mongo

import (
	"context"
	"fmt"
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Tenant   *mongo.Collection
	RBAC     *mongo.Collection
}

func NewMongoDB(logger *utils.Logger, ctx context.Context, uri string, dbname, tenantColl, rbacColl string) (*MongoDB, error) {
	logger.Info("connecting to mongo")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to mongo instance")
	}

	database := client.Database(dbname)
	tenant := database.Collection(tenantColl)
	rbac := database.Collection(rbacColl)

	logger.Info("creating indexes")
	if err := createTenantIndexes(logger, tenant); err != nil {
		return nil, errors.Wrap(err, "failed to create tenant indexes")
	}

	db := &MongoDB{
		Client:   client,
		Database: database,
		Tenant:   tenant,
		RBAC:     rbac,
	}

	return db, nil
}

func createTenantIndexes(logger *utils.Logger, collection *mongo.Collection) error {
	ctx := context.Background()

	logger.Info("creating indexes for tenant collection")
	indexSlice, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{
				"subdomain": 1,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{
				"primary_contacts.email": 1,
			},
			Options: options.Index().SetName("primary_contacts.email"),
		},
		{
			Keys: bson.M{
				"company_name": 1,
			},
			Options: options.Index().SetName("company_name"),
		},
		{
			Keys: bson.M{
				"subscription.plan": 1,
			},
			Options: options.Index().SetName("subscription.plan"),
		},
	})

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create indexes: %v", indexSlice))
	}

	logger.Infof("created indexes: %v", indexSlice)
	return nil
}
