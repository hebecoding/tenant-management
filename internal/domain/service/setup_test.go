package service_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/helpers"
	"github.com/hebecoding/tenant-management/infrastructure/config"
	"github.com/hebecoding/tenant-management/infrastructure/repositories/mongo"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	serv "github.com/hebecoding/tenant-management/internal/domain/service"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func TestMain(m *testing.M) {
	// configure test environment
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

		}
	}(client, context.Background())

	// create new collection for tenants
	collection := client.Database("test_tenants").Collection("tenants")
	mock.DB = collection

	logger.Info("Dropping existing test collections")
	if err := mock.DB.Drop(context.Background()); err != nil {
		logger.Fatal(err)
	}

	// create new tenant repository
	logger.Info("Creating new tenant repository")
	mock.Repo = mongo.NewTenantRepository(mock.DB, logger)

	// create test tenants
	logger.Info("Creating test tenants")
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-mock-data.json")
	if err != nil {
		logger.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var tenants []*entities.Tenant
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tenants); err != nil {
		logger.Fatal(err)
	}

	for _, tenant := range tenants {
		_, err := mock.DB.InsertOne(context.Background(), tenant)
		if err != nil {
			logger.Fatal(err)
		}
	}
	logger.Info("Successfully created test tenants")

	// create new tenant mock
	logger.Info("Creating new tenant mock")
	mock.Service = serv.NewTenantService(logger, mock.Repo)

	// run tests
	code := m.Run()

	os.Exit(code)
}
