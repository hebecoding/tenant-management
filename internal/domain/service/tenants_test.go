package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/infrastructure/repositories/mongo"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	serv "github.com/hebecoding/tenant-management/internal/domain/service"
	"github.com/hebecoding/tenant-management/tests"
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
	var testCases = []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
		CancelContext bool
	}{
		{
			Name:   "Service: Happy Path: Create a Tenant",
			Tenant: tests.CreateTenant(),
		},
	}

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

				err := mock.Service.CreateTenant(ctx, tt.Tenant)
				if expectedErr != err {
					assert.Equal(t, expectedErr.Error(), err.Error())
				}
			},
		)
	}
}

func TestTenantService_GetTenantByID(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
		CancelContext bool
	}{
		{
			Name: "Happy Path: Get a Tenant by ID",
			Tenant: func() *entities.Tenant {
				tenant := tests.CreateTenant()
				_ = mock.Service.CreateTenant(ctx, tenant)
				return tenant
			}(),
		},
		{
			Name: "Error Path: Get a Tenant by ID - Tenant not found",
			Tenant: func() *entities.Tenant {
				tenant := tests.CreateTenant()
				tenant.ID = "5f6a2b7e4c9e1f0001f1f8c5"
				return tenant
			}(),
			ExpectedError: "no tenant documents found",
		},
	}

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

				tenant, err := mock.Service.GetTenantByID(ctx, tt.Tenant.ID)
				if err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
					return
				}

				assert.Equal(t, tt.Tenant, tenant)
			},
		)
	}
}

func TestTenantService_UpdateTenant(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
	}{
		{
			Name: "Happy Path: Update a Tenant - Company Subscription Details",
			Tenant: func() *entities.Tenant {
				tenant := tests.CreateTenant()
				updatedTenant := &entities.Tenant{
					ID:        tenant.ID,
					Companies: []*entities.TenantCompanyDetails{tenant.Companies[0]},
				}

				_ = mock.Service.CreateTenant(ctx, tenant)

				newSubscriptionDetails := tests.GenerateSubscriptionDetails()
				newSubscriptionDetails.Active = true
				newSubscriptionDetails.StartDate.Truncate(time.Millisecond)
				newSubscriptionDetails.EndDate.Truncate(time.Millisecond)
				newSubscriptionDetails.NextBillingDate.Truncate(time.Millisecond)
				newSubscriptionDetails.LastPaymentDate.Truncate(time.Millisecond)

				updatedTenant.Companies[0].Subscriptions = []*entities.TenantSubscriptionDetails{newSubscriptionDetails}

				return updatedTenant
			}(),
		},
	}

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

				tenant := &entities.Tenant{}

				err := mock.Service.UpdateTenant(ctx, tt.Tenant.ID, tenant)
				if err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
					return
				}
			},
		)
	}
}

func TestTenantService_DeleteTenant(t *testing.T) {
	var testCases = []struct {
		Name          string
		TenantID      string
		ExpectedError string
	}{
		{
			Name: "Happy Path: Delete a Tenant",
			TenantID: func() string {
				tenant := tests.CreateTenant()
				_ = mock.Service.CreateTenant(ctx, tenant)
				return tenant.ID
			}(),
		},
	}

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

				err := mock.Service.DeleteTenant(ctx, tt.TenantID)
				if err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
					return
				}

				tenant, _ := mock.Service.GetTenantByID(ctx, tt.TenantID)
				assert.False(t, tenant.IsActive)
			},
		)
	}
}

func TestTenantService_GetTenants(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenants       []*entities.Tenant
		ExpectedError string
	}{
		{
			Name: "Happy Path: Get all Tenants",
			Tenants: func() []*entities.Tenant {
				tenants := []*entities.Tenant{
					tests.CreateTenant(),
					tests.CreateTenant(),
					tests.CreateTenant(),
				}

				for _, tenant := range tenants {
					_ = mock.Service.CreateTenant(ctx, tenant)
				}

				return tenants
			}(),
		},
	}

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

				tenants, err := mock.Service.GetTenants(ctx)
				if err != nil {
					assert.EqualError(t, err, expectedErr.Error())
					return
				}

				assert.Len(t, tenants, len(tt.Tenants))
			},
		)
	}
}

func TestTenantService_GetTenantCompanies(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
	}{
		{
			Name: "Happy Path: Get Tenant Companies",
			Tenant: func() *entities.Tenant {
				tenant := tests.CreateTenant()
				updatedTenant := &entities.Tenant{ID: tenant.ID, Companies: tenant.Companies}
				_ = mock.Service.CreateTenant(ctx, tenant)

				return updatedTenant
			}(),
		},
	}

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

				tenantCompanies, err := mock.Service.GetTenantCompanies(context.Background(), tt.Tenant.ID)
				if err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
					return
				}

				assert.EqualValues(t, tt.Tenant.Companies, tenantCompanies)
			},
		)
	}
}

func TestTenantService_GetTenantCompaniesSubscriptions(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
	}{
		{
			Name: "Happy Path: Get Tenant Companies Subscriptions",
			Tenant: func() *entities.Tenant {
				tenant := tests.CreateTenant()
				tenantVals := &entities.Tenant{ID: tenant.ID, Companies: tenant.Companies}

				_ = mock.Service.CreateTenant(context.Background(), tenant)

				return tenantVals
			}(),
		},
	}

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

				subscriptions, err := mock.Service.GetTenantCompaniesSubscriptions(ctx, tt.Tenant.ID)
				if err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
					return
				}

				testSubscriptions := []*entities.TenantSubscriptionDetails{}

				for _, val := range tt.Tenant.Companies {
					testSubscriptions = append(testSubscriptions, val.Subscriptions...)
				}

				assert.EqualValues(t, testSubscriptions, subscriptions)
			},
		)
	}
}

func TestTenantService_UpdateTenantSubscription(t *testing.T) {
	var testCases = []struct {
		Name          string
		Tenant        *entities.Tenant
		ExpectedError string
	}{
		{
			Name: "Happy Path: Update Tenant Subscription",
			Tenant: func() *entities.Tenant {
				tenant := tests.CreateTenant()
				tenant.IsActive = true
				tenant.Companies = tenant.Companies[:1]

				_ = mock.Service.CreateTenant(context.Background(), tenant)

				newSub := tests.GenerateSubscriptionDetails()
				newSub.ID = tenant.Companies[0].Subscriptions[0].ID

				updatedTenantVals := &entities.Tenant{ID: tenant.ID, Companies: tenant.Companies, IsActive: true}
				updatedTenantVals.Companies[0].Subscriptions = []*entities.TenantSubscriptionDetails{newSub}

				return updatedTenantVals
			}(),
		},
	}

	for _, tt := range testCases {
		defer t.Run(
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

				err := mock.Service.UpdateTenantSubscription(
					ctx,
					tt.Tenant.ID,
					tt.Tenant.Companies[0].Subscriptions[0],
				)

				if err != nil {
					assert.Equal(t, expectedErr.Error(), err.Error())
					return
				}

				subscriptions, err := mock.Service.GetTenantCompaniesSubscriptions(ctx, tt.Tenant.ID)
				if err != nil {
					assert.NoError(t, err)
				}

				assert.EqualValues(t, tt.Tenant.Companies[0].Subscriptions, subscriptions)
			},
		)
	}

}
