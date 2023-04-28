package service

import (
	"context"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
)

type TenantService interface {
	CreateTenant(ctx context.Context, tenant *entities.Tenant) (*entities.Tenant, error)
	GetTenantByID(ctx context.Context, id string) (*entities.Tenant, error)
	UpdateTenant(ctx context.Context, tenant *entities.Tenant) (*entities.Tenant, error)
	DeleteTenant(ctx context.Context, id string) error
	ListTenants(ctx context.Context, offset, limit int64) ([]*entities.Tenant, error)
	GetTenantSubscription(ctx context.Context, id string) (*entities.TenantSubscriptionDetails, error)
	UpdateTenantSubscription(ctx context.Context, id string, subscription *entities.TenantSubscriptionDetails) error
	GetTenantCompany(ctx context.Context, id string) (*entities.TenantCompanyDetails, error)
	UpdateTenantCompany(ctx context.Context, id string, company *entities.TenantCompanyDetails) error
	GetTenantPaymentDetails(ctx context.Context, id string) (*entities.TenantPaymentDetails, error)
	UpdateTenantPaymentDetails(ctx context.Context, id string, paymentDetails *entities.TenantPaymentDetails) error
	UpdateTenantFields(ctx context.Context, id string, fieldsToUpdate map[string]interface{}) (*entities.Tenant, error)
}
