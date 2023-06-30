package mongo_test

import (
	"context"
	"testing"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/infrastructure/repositories/mongo"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	"github.com/hebecoding/tenant-management/tests"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

type TestTenantRepository struct {
	DB   *mgo.Collection
	Repo *mongo.TenantRepository
}

var storage = &TestTenantRepository{}
var logger utils.LoggerInterface
var ctx = context.Background()

func TestTenantRepository_Create(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
		CancelContext bool
	}{
		{
			Name:          "Happy Path: Create Tenant successfully",
			Tenant:        tests.CreateTenant(),
			ExpectedError: "",
			CancelContext: false,
		},
	}

	// run test cases
	for _, tt := range testCases {
		t.Run(
			tt.Name, func(t *testing.T) {
				defer func() {
					err := dropTestCollections()
					if err != nil {
						logger.Error(err)
					}
				}()

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
				if gotErr := storage.Repo.CreateTenant(ctx, tt.Tenant); gotErr != expectedErr {
					assert.Equal(t, expectedErr, gotErr)
				}
			},
		)
	}
}

func TestTenantRepository_GetTenants(t *testing.T) {
	var testCases = []struct {
		Name          string
		ExpectedError string
		Tenants       []*entities.Tenant
	}{
		{
			Name:          "Happy Path: Get Tenants successfully",
			ExpectedError: "",
			Tenants: func() []*entities.Tenant {
				tenantSlice := []*entities.Tenant{}

				for i := 0; i < 10; i++ {
					tenant := tests.CreateTenant()
					_ = storage.Repo.CreateTenant(ctx, tenant)
					tenantSlice = append(tenantSlice, tenant)
				}

				return tenantSlice
			}(),
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run test cases
	for _, tt := range testCases {
		t.Run(
			tt.Name, func(t *testing.T) {
				defer func() {
					err := dropTestCollections()
					if err != nil {
						logger.Error(err)
					}
				}()

				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				tenants, gotErr := storage.Repo.GetTenants(ctx)
				if gotErr != expectedErr {
					assert.Fail(t, "expected error", expectedErr, "got error", gotErr)
				}
				assert.Len(t, tenants, len(tt.Tenants))
			},
		)
	}
}

func TestTenantRepository_GetTenantByID(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
	}{
		{
			Name: "Happy Path: Get Tenant by ID successfully",
			Tenant: func() *entities.Tenant {
				tenant := tests.CreateTenant()
				tenant.IsActive = true
				_ = storage.Repo.CreateTenant(ctx, tenant)
				return tenant
			}(),
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run test cases
	for _, tt := range testCases {
		t.Run(
			tt.Name, func(t *testing.T) {
				defer func() {
					err := dropTestCollections()
					if err != nil {
						logger.Error(err)
					}
				}()

				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				tenant, gotErr := storage.Repo.GetTenantByID(ctx, tt.Tenant.ID)
				if gotErr != nil {
					assert.EqualError(t, gotErr, expectedErr.Error())
				}

				assert.EqualValues(t, tt.Tenant, tenant)

			},
		)
	}
}

func TestTenantRepository_UpdateTenant(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenant        entities.Tenant
		UpdatedTenant entities.Tenant
		ExpectedError string
	}{
		{
			Name: "Happy Path: Update Tenant successfully",
			Tenant: func() entities.Tenant {
				tenant := tests.CreateTenant()
				_ = storage.Repo.CreateTenant(ctx, tenant)

				payDetailsID := tenant.PaymentDetails[0].ID
				newPaymentDetails := tests.GeneratePaymentDetails()
				newPaymentDetails.ID = payDetailsID

				newTenant := entities.Tenant{}
				newTenant.ID = tenant.ID
				newTenant.PaymentDetails = append(newTenant.PaymentDetails, newPaymentDetails)
				return newTenant
			}(),
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run test cases
	for _, tt := range testCases {
		t.Run(
			tt.Name, func(t *testing.T) {
				defer func() {
					err := dropTestCollections()
					if err != nil {
						logger.Error(err)
					}
				}()

				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				gotErr := storage.Repo.UpdateTenant(ctx, &tt.Tenant)
				if gotErr != expectedErr {
					assert.Fail(t, "expected error", expectedErr, "got error", gotErr)
				}
			},
		)
	}
}

func TestTenantRepository_DeleteTenant(t *testing.T) {
	var testCases = []struct {
		Name          string
		TenantID      string
		ExpectedError string
	}{
		{
			Name: "Happy Path: Delete Tenant successfully",
			TenantID: func() string {
				tenant := tests.CreateTenant()
				_ = storage.Repo.CreateTenant(ctx, tenant)
				return tenant.ID
			}(),
			ExpectedError: "",
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run test cases
	for _, tt := range testCases {
		t.Run(
			tt.Name, func(t *testing.T) {
				defer func() {
					err := dropTestCollections()
					if err != nil {
						logger.Error(err)
					}
				}()

				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				if gotErr := storage.Repo.DeleteTenant(ctx, tt.TenantID); gotErr != expectedErr {
					assert.ErrorContains(t, gotErr, expectedErr.Error())
				}

				tenant, err := storage.Repo.GetTenantByID(ctx, tt.TenantID)
				assert.NoError(t, err)

				assert.False(t, tenant.IsActive)
			},
		)
	}
}

func TestTenantRepository_SearchTenant(t *testing.T) {

	tenant := tests.CreateTenant()
	_ = storage.Repo.CreateTenant(ctx, tenant)

	var testCases = []struct {
		Name          string
		TenantID      string
		SearchParams  map[string]any
		ExpectedError string
	}{
		{
			Name:     "Happy Path: Search Tenant successfully",
			TenantID: tenant.ID,
			SearchParams: map[string]any{
				"payment_details._id": tenant.PaymentDetails[0].ID,
			},
			ExpectedError: "",
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run test cases
	for _, tt := range testCases {
		t.Run(
			tt.Name, func(t *testing.T) {
				defer func() {
					err := dropTestCollections()
					if err != nil {
						logger.Error(err)
					}
				}()

				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				expectedTenant, err := storage.Repo.GetTenantByID(ctx, tt.TenantID)
				assert.Nil(t, err)

				gotTenant, gotErr := storage.Repo.SearchTenant(ctx, tt.SearchParams)
				if gotErr != expectedErr {
					assert.Nil(t, gotErr)
				}
				assert.EqualValues(t, expectedTenant, gotTenant)
			},
		)
	}
}

func TestTenantRepository_SearchTenants(t *testing.T) {
	var testCases = []struct {
		Name          string
		SearchParams  map[string]any
		ExpectedError string
		Tenants       []*entities.Tenant
	}{
		{
			Name: "Happy Path: Search Tenants successfully",
			SearchParams: map[string]any{
				"is_active": true,
			},
			Tenants: func() []*entities.Tenant {
				var tenants []*entities.Tenant
				for i := 0; i < 5; i++ {
					tenant := tests.CreateTenant()
					tenant.IsActive = true
					_ = storage.Repo.CreateTenant(ctx, tenant)
					tenants = append(tenants, tenant)
				}
				return tenants
			}(),
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run test cases
	for _, tt := range testCases {
		t.Run(
			tt.Name, func(t *testing.T) {
				defer func() {
					err := dropTestCollections()
					if err != nil {
						logger.Error(err)
					}
				}()

				var expectedErr error
				if tt.ExpectedError != "" {
					expectedErr = errors.New(tt.ExpectedError)
				}

				tenants, gotErr := storage.Repo.SearchTenants(ctx, tt.SearchParams)
				if gotErr != expectedErr {
					assert.NoError(t, gotErr)
				}

				assert.Len(t, tenants, len(tt.Tenants))
			},
		)
	}
}
