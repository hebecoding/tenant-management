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

	if err := client.Ping(context.Background(), nil); err != nil {
		logger.Fatal(errors.Wrap(err, "error pinging mongo test database"))
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

	// create test tenants
	logger.Info("Creating test tenants")
	file, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-mock-data.json")
	if err != nil {
		logger.Fatal(err)
	}
	defer file.Close()

	var tenants []*entities.Tenant
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tenants); err != nil {
		logger.Fatal(err)
	}

	for _, tenant := range tenants {
		_, err := storage.DB.InsertOne(context.Background(), tenant)
		if err != nil {
			logger.Fatal(err)
		}
	}
	logger.Info("Successfully created test tenants")

	// run tests
	code := m.Run()

	// drop test collections
	logger.Info("Dropping test collections")
	if err := storage.DB.Drop(context.Background()); err != nil {
		logger.Fatal(err)
	}

	os.Exit(code)

}

func TestTenantRepository_Create(t *testing.T) {
	// read in test data from file
	file, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-create.json")
	if err != nil {
		assert.Nil(t, err)
	}

	logger.Infoln("Successfully read in test data from file")
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
		assert.Nil(t, err)
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

				// when
				if gotErr := storage.Repo.Create(ctx, tt.Tenant); gotErr != expectedErr {
					assert.Equal(t, expectedErr, gotErr)
				}
			},
		)
	}

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
	testFile, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-get-by-id.json")
	if err != nil {
		assert.Nil(t, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tests []struct {
		Name          string
		TenantID      string
		ExpectedError string
	}

	// decode test data
	decoder := json.NewDecoder(testFile)
	_ = decoder.Decode(&tests)

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
	testFile, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-update.json")
	if err != nil {
		assert.Nil(t, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tests []struct {
		Name          string
		TenantID      string
		UpdatedValues map[string]interface{}
		ExpectedError string
	}

	// decode test data
	decoder := json.NewDecoder(testFile)
	_ = decoder.Decode(&tests)

	// run test cases
	for _, tt := range tests {
		newTenant := &entities.Tenant{}
		var testValues struct {
			PaymentDetails []*entities.TenantPaymentDetails `json:"payment_details"`
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

func TestTenantRepository_DeleteTenant(t *testing.T) {
	// read in test data from testData
	testFile, err := readInJSONTestDataFile("../../../tests/test-data/storage/tenant-delete.json")
	assert.Nil(t, err)
	defer testFile.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tests []struct {
		Name          string
		TenantID      string
		ExpectedError string
	}

	// decode test data
	decoder := json.NewDecoder(testFile)
	_ = decoder.Decode(&tests)

	// run test cases
	for _, tt := range tests {
		t.Run(
			tt.Name, func(t *testing.T) {
				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				if gotErr := storage.Repo.DeleteTenant(ctx, tt.TenantID); gotErr != expectedErr {
					assert.ErrorContains(t, gotErr, expectedErr.Error())
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
