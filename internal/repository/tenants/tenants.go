package tenants

import (
	"context"
	"github.com/hebecoding/tenant-management/internal/domain/tenants"
	"go.mongodb.org/mongo-driver/mongo"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *tenants.Tenant) error
	GetByID(ctx context.Context, id string) (*tenants.Tenant, error)
	Update(ctx context.Context, id string, update *tenants.Tenant) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*tenants.Tenant, error)
	GetSubscriptionDetails(ctx context.Context, id string) (*tenants.TenantSubscriptionDetails, error)
	UpdateSubscriptionDetails(ctx context.Context, id string, update *tenants.TenantSubscriptionDetails) error
	GetCompanyDetails(ctx context.Context, id string) (*tenants.TenantCompanyDetails, error)
	UpdateCompanyDetails(ctx context.Context, id string, update *tenants.TenantCompanyDetails) error
	GetPaymentDetails(ctx context.Context, id string) (*tenants.TenantPaymentDetails, error)
	UpdatePaymentDetails(ctx context.Context, id string, update *tenants.TenantPaymentDetails) error
}

type TenantRepo struct {
	collection *mongo.Collection
}

func NewTenantRepository() *TenantRepo {
	return &TenantRepo{}
}
