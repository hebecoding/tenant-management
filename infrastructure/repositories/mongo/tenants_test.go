package mongo_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/infrastructure/config"
	"github.com/hebecoding/tenant-management/infrastructure/repositories/mongo"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestTenantRepository struct {
	DB   *mgo.Collection
	Repo *mongo.TenantRepository
}

var storage = &TestTenantRepository{}
var logger *utils.Logger

func TestMain(m *testing.M) {
	// initialize logger
	logger = utils.NewLogger()
	// read in config
	if err := config.ReadInConfig(logger); err != nil {
		logger.Fatal(err)
	}

	// connect to mongo test database
	logger.Info("Connecting to mongo test database")
	client, err := mgo.Connect(context.Background(), options.Client().ApplyURI(config.Config.DB.URL))
	if err != nil {
		logger.Infoln("error connecting to mongo test database")
		logger.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	// create new test collection for tenants
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

func TestTenantRepository_Create(t *testing.T) {
	// read in test data from file
	file, err := readInJSONTestDataFile("test-data/tenant-create.json")
	if err != nil {
		assert.Nil(t, err)
	}

	var tests []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
		CancelContext bool
	}

	// decode test data
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tests); err != nil {
		assert.Nil(t, err)
		logger.Error("error decoding test data: ", err)
	}

	// run test cases
	for _, tt := range tests {
		t.Run(
			tt.Name, func(t *testing.T) {
				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				if tt.CancelContext {
					expectedErr = context.Canceled
					cancel()
				}

				if gotErr := storage.Repo.Create(ctx, tt.Tenant); gotErr != expectedErr {
					assert.Equal(t, expectedErr, gotErr)
				}
			},
		)
	}

}

func TestTenantRepository_GetTenants(t *testing.T) {
	// read in test data from file
	file, err := readInJSONTestDataFile("test-data/tenant-get-tenants.json")
	if err != nil {
		assert.Nil(t, err)
	}

	defer file.Close()

	var tests []struct {
		Name          string
		Tenants       []*entities.Tenant
		ExpectedError string
	}

	// decode test data
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tests); err != nil {
		assert.Nil(t, err)
		logger.Error("error decoding test data: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create tenants
	for _, tt := range tests {
		for _, tenant := range tt.Tenants {
			if gotErr := storage.Repo.Create(ctx, tenant); gotErr != nil {
				assert.Nil(t, gotErr)
			}
		}
	}

	// run test cases
	for _, tt := range tests {
		t.Run(
			tt.Name, func(t *testing.T) {
				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				tenants, gotErr := storage.Repo.GetTenants(ctx)
				if gotErr != expectedErr {
					assert.Equal(t, expectedErr, gotErr)
				}
				fmt.Println(tenants)
			},
		)
	}

}

func readInJSONTestDataFile(path string) (*os.File, error) {
	// read in test data from file
	logger.Info("Reading in test data from file")
	file, err := os.Open(path)
	if err != nil {
		logger.Error("error opening file in path: ", path)
		return nil, errors.Wrap(err, "error opening test data file")
	}
	logger.Info("Successfully read in test data from file")
	return file, nil
}
