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
	"github.com/stretchr/testify/assert"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestTenantService struct {
	Service *serv.TenantService
	DB      *mgo.Collection
	Repo    *mongo.TenantRepository
}

var (
	logger  utils.LoggerInterface
	ctx     = context.Background()
	service = &TestTenantService{}
)

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
	defer container.Terminate(ctx)

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

	defer client.Disconnect(context.Background())

	// create new collection for tenants
	collection := client.Database("test_tenants").Collection("tenants")
	service.DB = collection

	logger.Info("Dropping existing test collections")
	if err := service.DB.Drop(context.Background()); err != nil {
		logger.Fatal(err)
	}

	// create new tenant repository
	logger.Info("Creating new tenant repository")
	service.Repo = mongo.NewTenantRepository(service.DB, logger)

	// create test tenants
	logger.Info("Creating test tenants")
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-mock-data.json")
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
		_, err := service.DB.InsertOne(context.Background(), tenant)
		if err != nil {
			logger.Fatal(err)
		}
	}
	logger.Info("Successfully created test tenants")

	// create new tenant service
	logger.Info("Creating new tenant service")
	service.Service = serv.NewTenantService(logger, service.Repo)

	// run tests
	code := m.Run()

	os.Exit(code)
}

func TestTenantService_CreateTenant(t *testing.T) {
	// read in test data
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-create.json")
	assert.NoError(t, err)
	defer file.Close()

	var tests []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
		CancelContext bool
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tests); err != nil {
		logger.Error(err)
	}

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

				err = service.Service.CreateTenant(ctx, tt.Tenant)
				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}
}

func TestTenantService_GetTenantByID(t *testing.T) {
	// read in test data
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-get-by-id.json")
	assert.NoError(t, err)
	defer file.Close()

	var tests []struct {
		Name          string
		TenantID      string
		ExpectedError string
		CancelContext bool
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tests); err != nil {
		logger.Error(err)
	}

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

				_, err = service.Service.GetTenantByID(ctx, tt.TenantID)

				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}
}

func TestTenantService_UpdateTenant(t *testing.T) {
	// read in test data
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-update.json")
	if err != nil {
		logger.Error(err)
	}
	defer file.Close()

	var tests []struct {
		Name          string
		TenantID      string
		UpdatedValues map[string]any
		ExpectedError string
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tests); err != nil {
		logger.Error(err)
	}

	for _, tt := range tests {
		t.Run(
			tt.Name, func(t *testing.T) {

				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				tenant := &entities.Tenant{}
				vals, err := json.Marshal(tt.UpdatedValues)
				if err != nil {
					logger.Error(err)
				}

				if err := json.Unmarshal(vals, tenant); err != nil {
					logger.Error(err)
				}

				err = service.Service.UpdateTenant(ctx, tt.TenantID, tenant)
				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}
}

func TestTenantService_DeleteTenant(t *testing.T) {
	// read in test data
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-delete.json")
	if err != nil {
		logger.Error(err)
	}
	defer file.Close()

	var tests []struct {
		Name          string
		TenantID      string
		ExpectedError string
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tests); err != nil {
		logger.Error(err)
	}

	for _, tt := range tests {
		t.Run(
			tt.Name, func(t *testing.T) {
				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				err = service.Service.DeleteTenant(ctx, tt.TenantID)
				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}
}

func TestTenantService_GetTenants(t *testing.T) {
	// read in test data
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-get-tenants.json")
	if err != nil {
		logger.Error(err)
	}
	defer file.Close()

	var tests []struct {
		Name          string
		ExpectedError string
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tests); err != nil {
		logger.Error(err)
	}

	for _, tt := range tests {
		t.Run(
			tt.Name, func(t *testing.T) {
				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				_, err = service.Service.GetTenants(ctx)
				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}
}

func TestTenantService_GetTenantCompanies(t *testing.T) {
	// read in test data
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-get-companies.json")
	if err != nil {
		logger.Error(err)
	}
	defer file.Close()

	var tests []struct {
		Name          string
		TenantID      string
		ExpectedError string
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tests); err != nil {
		logger.Error(err)
	}

	for _, tt := range tests {
		t.Run(
			tt.Name, func(t *testing.T) {
				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				_, err = service.Service.GetTenantCompanies(ctx, tt.TenantID)
				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}
}
