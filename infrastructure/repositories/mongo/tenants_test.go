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

	// drop test collections
	logger.Info("Dropping test collections")
	if err := storage.DB.Drop(context.Background()); err != nil {
		logger.Fatal(err)
	}

	os.Exit(code)

}

func TestTenantRepository_GetTenants(t *testing.T) {
	// read in test data from file
	file, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-get-tenants.json")
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

func TestTenantRepository_GetTenantByID(t *testing.T) {
	// read in test data from testData
	testData, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-mock-data.json")
	if err != nil {
		assert.Nil(t, err)
	}

	testFile, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-get-by-id.json")
	if err != nil {
		assert.Nil(t, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer testData.Close()
	defer cancel()

	var testTenants []*entities.Tenant
	var tests []struct {
		Name          string
		TenantID      string
		ExpectedError string
	}

	// decode test data
	decoder := json.NewDecoder(testData)
	_ = decoder.Decode(&testTenants)
	decoder = json.NewDecoder(testFile)
	_ = decoder.Decode(&tests)

	// create tenants+
	for _, tt := range testTenants {
		if gotErr := storage.Repo.Create(ctx, tt); gotErr != nil {
			assert.Nil(t, gotErr)
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

				_, gotErr := storage.Repo.GetTenantByID(ctx, tt.TenantID)
				if gotErr != expectedErr {
					assert.ErrorContains(t, gotErr, expectedErr.Error())
				}
			},
		)
	}

}

func TestTenantRepository_UpdateTenant(t *testing.T) {
	// read in test data from testData
	testData, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-mock-data.json")
	if err != nil {
		assert.Nil(t, err)
	}

	testFile, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-update.json")
	if err != nil {
		assert.Nil(t, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer testData.Close()
	defer cancel()

	var testTenants []*entities.Tenant
	var tests []struct {
		Name          string
		TenantID      string
		UpdatedValues map[string]interface{}
		ExpectedError string
	}

	// decode test data
	decoder := json.NewDecoder(testData)
	_ = decoder.Decode(&testTenants)
	decoder = json.NewDecoder(testFile)
	_ = decoder.Decode(&tests)

	// create tenants+
	for _, tt := range testTenants {
		if gotErr := storage.Repo.Create(ctx, tt); gotErr != nil {
			assert.Nil(t, gotErr)
		}
	}

	// run test cases
	for _, tt := range tests {
		newTenant := &entities.Tenant{}
		var testValues struct {
			PaymentDetails *entities.TenantPaymentDetails `json:"payment_details"`
		}

		vals, err := json.Marshal(tt.UpdatedValues)
		assert.NoError(t, err)

		if err := json.Unmarshal(vals, &testValues); err != nil {
			assert.Nil(t, err)
		}

		t.Run(
			tt.Name, func(t *testing.T) {
				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				tenant, gotErr := storage.Repo.GetTenantByID(ctx, tt.TenantID)
				if gotErr != expectedErr {
					assert.ErrorContains(t, gotErr, expectedErr.Error())
				}

				switch {
				case testValues.PaymentDetails != nil:
					newTenant.ID = tenant.ID
					newTenant.PaymentDetails = testValues.PaymentDetails

					if gotErr := storage.Repo.UpdateTenant(ctx, newTenant); gotErr != expectedErr {
						assert.ErrorContains(t, gotErr, expectedErr.Error())
					}

					tenant, err = storage.Repo.GetTenantByID(ctx, tt.TenantID)
					assert.Nil(t, err)

					assert.EqualValues(t, newTenant.PaymentDetails, tenant.PaymentDetails)
				}

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
