package mongo_test

import (
	"os"
	"testing"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/infrastructure/config"
	"github.com/hebecoding/tenant-management/infrastructure/repositories/mongo"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func TestMain(m *testing.M) {
	// initialize logger
	logger = utils.NewLogger()
	// read in config
	if err := config.ReadInConfig(logger); err != nil {
		logger.Fatal(err)
	}

	// connect to mongo test database
	logger.Info("Connecting to mongo test database")

	// connect to mongo testcontainers
	container, err := NewMongoDBTestContainer(ctx)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(container *MongoDBTestContainer, ctx context.Context) {
		err := container.Terminate(ctx)
		if err != nil {
			logger.Fatal(err)
		}
	}(container, ctx)

	endpoint, err := container.Endpoint(ctx, "mongodb")
	if err != nil {
		logger.Fatal(err)
	}

	client, err := mgo.Connect(ctx, options.Client().ApplyURI(endpoint))
	if err != nil {
		logger.Info("error connecting to mongo test database")
		logger.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		logger.Fatal(errors.Wrap(err, "error pinging mongo test database"))
	}

	defer func(client *mgo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			logger.Fatal(err)
		}
	}(client, context.Background())

	// create new collection for tenants
	collection := client.Database("test_tenants").Collection("tenants")
	storage.DB = collection

	logger.Info("Dropping existing test collections")
	if err := storage.DB.Drop(context.Background()); err != nil {
		logger.Fatal(err)
	}

	// create new tenant repository
	logger.Info("Creating new tenant repository")
	storage.Repo = mongo.NewTenantRepository(storage.DB, logger)

	// run tests
	code := m.Run()

	os.Exit(code)

}

type MongoDBTestContainer struct {
	testcontainers.Container
}

func NewMongoDBTestContainer(ctx context.Context) (*MongoDBTestContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:6.0.5",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForLog("Waiting for connections"),
			wait.ForListeningPort("27017/tcp"),
		),
	}

	container, err := testcontainers.GenericContainer(
		ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		return nil, err
	}

	return &MongoDBTestContainer{
		Container: container,
	}, nil
}

func dropTestCollections() error {
	// drop existing collections
	logger.Info("Dropping existing test collections")
	if err := storage.DB.Drop(context.Background()); err != nil {
		logger.Fatal(err)
	}

	return nil
}
