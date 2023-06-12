package service

import (
	"context"

	"github.com/hebecoding/tenant-management/internal/domain/entities"
)

type TenantService interface {
	CreateTenant(ctx context.Context, tenant *entities.Tenant) error
	DeleteTenant(ctx context.Context, id string) error
	GetTenantByID(ctx context.Context, id string) (*entities.Tenant, error)
	GetTenantCompanies(ctx context.Context, id string) ([]*entities.TenantCompanyDetails, error)
	GetTenantCompanyByID(ctx context.Context, id string, companyID string) (*entities.TenantCompanyDetails, error)
	GetTenantPaymentDetails(ctx context.Context, id string) ([]*entities.TenantPaymentDetails, error)
	GetTenantByPaymentID(ctx context.Context, paymentID string) (*entities.Tenant, error)
	GetTenantSubscription(ctx context.Context, id string) (*entities.TenantSubscriptionDetails, error)
	GetTenants(ctx context.Context) ([]*entities.Tenant, error)
	UpdateTenant(ctx context.Context, id string, tenant *entities.Tenant) error
	UpdateTenantCompany(ctx context.Context, id string, company *entities.TenantCompanyDetails) error
	UpdateTenantSubscription(ctx context.Context, id string, subscription *entities.TenantSubscriptionDetails) error
	UpdateTenantPaymentDetails(ctx context.Context, id string, paymentDetails *entities.TenantPaymentDetails) error
}
