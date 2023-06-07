package repository

import (
	"context"

	"github.com/hebecoding/tenant-management/internal/domain/entities"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *entities.Tenant) error
	DeleteTenant(ctx context.Context, id string) error
	GetTenantByID(ctx context.Context, id string) (*entities.Tenant, error)
	GetTenants(ctx context.Context) ([]*entities.Tenant, error)
	UpdateTenant(ctx context.Context, tenant *entities.Tenant) error
}
