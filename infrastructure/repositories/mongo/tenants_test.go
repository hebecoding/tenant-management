package mongo_test

import (
	"context"
	"encoding/json"
	"log"
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

var TestRepo = &TestTenantRepository{}
var Logger *utils.Logger

func TestMain(m *testing.M) {
	// initialize logger
	Logger = utils.NewLogger()
	// read in config
	if err := config.ReadInConfig(Logger); err != nil {
		Logger.Fatal(err)
	}

	// connect to mongo test database
	Logger.Info("Connecting to mongo test database")
	client, err := mgo.Connect(context.Background(), options.Client().ApplyURI(config.Config.DB.Url))
	if err != nil {
		Logger.Infoln("error connecting to mongo test database")
		Logger.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	// create new test collection for tenants
	collection := client.Database("test_tenants").Collection("tenants")
	TestRepo.DB = collection

	Logger.Info("Dropping existing test collections")
	if err := TestRepo.DB.Drop(context.Background()); err != nil {
		Logger.Fatal(err)
	}

	// create new tenant repository
	Logger.Info("Creating new tenant repository")
	TestRepo.Repo = mongo.NewTenantRepository(TestRepo.DB, Logger)

	// run tests
	code := m.Run()

	os.Exit(code)

}

func TestTenantRepository_Create(t *testing.T) {
	// read in test data from file
	log.Println("Reading in test data from file")
	file, err := os.Open("test-data/tenant-create.json")
	if err != nil {
		t.Fatalf("error opening test data file: %v", err)
	}
	log.Println("Successfully read in test data from file")
	defer file.Close()

	var tests []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
		CancelContext bool
	}

	// decode test data
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tests); err != nil {
		t.Fatalf("error decoding test data: %v", err)
	}

	// run test cases
	for _, tt := range tests {
		t.Run(
			tt.Name, func(t *testing.T) {
				log.Println("Running test case: ", tt.Name)
				log.Println("Test Data: ", tt.Tenant)

				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				ctx, cancel := context.WithCancel(context.Background())
				if tt.CancelContext {
					expectedErr = context.Canceled
					cancel()
				}

				// when
				if gotErr := TestRepo.Repo.Create(ctx, tt.Tenant); gotErr != expectedErr {
					assert.Equal(t, expectedErr, gotErr)
				}
			},
		)
	}

}
