package service_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/helpers"
	"github.com/hebecoding/tenant-management/infrastructure/repositories/mongo"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	serv "github.com/hebecoding/tenant-management/internal/domain/service"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

type TestTenantService struct {
	Service *serv.TenantService
	DB      *mgo.Collection
	Repo    *mongo.TenantRepository
}

var (
	logger utils.LoggerInterface
	ctx    = context.Background()
	mock   = &TestTenantService{}
)

func TestTenantService_CreateTenant(t *testing.T) {
	// read in test data
	file, err := helpers.ReadInJSONTestDataFile(logger, "../../../tests/test-data/storage/tenant-create.json")
	assert.NoError(t, err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

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

				err = mock.Service.CreateTenant(ctx, tt.Tenant)
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

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

				_, err = mock.Service.GetTenantByID(ctx, tt.TenantID)

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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

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

				err = mock.Service.UpdateTenant(ctx, tt.TenantID, tenant)
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

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

				err = mock.Service.DeleteTenant(ctx, tt.TenantID)
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

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

				_, err = mock.Service.GetTenants(ctx)
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

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

				_, err = mock.Service.GetTenantCompanies(ctx, tt.TenantID)
				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}
}

func TestTenantService_GetTenantCompaniesSubscriptions(t *testing.T) {
	// read in test data
	file, err := helpers.ReadInJSONTestDataFile(
		logger, "../../../tests/test-data/storage/tenant-get-company-subscriptions.json",
	)
	if err != nil {
		assert.NoError(t, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			assert.NoError(t, err)
		}
	}(file)

	var tests []struct {
		Name          string
		TenantID      string
		ExpectedError string
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tests); err != nil {
		assert.NoError(t, err)
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

				_, err = mock.Service.GetTenantCompaniesSubscriptions(ctx, tt.TenantID)
				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}
}

func TestTenantService_UpdateTenantSubscription(t *testing.T) {
	// read in tenant data
	file, err := helpers.
		ReadInJSONTestDataFile(
			logger,
			"../../../tests/test-data/storage/tenant-update-subscription.json",
		)

	if err != nil {
		assert.NoError(t, err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			assert.NoError(t, err)
		}
	}(file)

	var tests []struct {
		Name          string
		TenantID      string
		Tenant        *entities.Tenant
		UpdatedValues map[string]*entities.TenantSubscriptionDetails
		ExpectedError string
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&tests); err != nil {
		assert.NoError(t, err)
	}

	// create test tenant
	if err = mock.Service.CreateTenant(context.Background(), tests[0].Tenant); err != nil {
		assert.NoError(t, err)
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

				err = mock.Service.UpdateTenantSubscription(
					ctx,
					tt.TenantID,
					tt.UpdatedValues["subscriptionDetails"],
				)

				if expectedErr != nil && err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
				} else {
					assert.Equal(t, expectedErr, err)
				}
			},
		)
	}

}
