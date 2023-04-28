package repository

import (
	"context"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *entities.Tenant) error
	GetByID(ctx context.Context, id string) (*entities.Tenant, error)
	Update(ctx context.Context, id string, update *entities.Tenant) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*entities.Tenant, error)
	GetSubscriptionDetails(ctx context.Context, id string) (*entities.TenantSubscriptionDetails, error)
	UpdateSubscriptionDetails(ctx context.Context, id string, update *entities.TenantSubscriptionDetails) error
	GetCompanyDetails(ctx context.Context, id string) (*entities.TenantCompanyDetails, error)
	UpdateCompanyDetails(ctx context.Context, id string, update *entities.TenantCompanyDetails) error
	GetPaymentDetails(ctx context.Context, id string) (*entities.TenantPaymentDetails, error)
	UpdatePaymentDetails(ctx context.Context, id string, update *entities.TenantPaymentDetails) error
}

type TenantRepo struct {
	db *mongo.Collection
}

func NewTenantRepository() *TenantRepo {
	return &TenantRepo{}
}
