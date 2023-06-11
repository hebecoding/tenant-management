package repository

import (
	"context"

	"github.com/hebecoding/tenant-management/internal/domain/entities"
)

type TenantRepository interface {
	CreateTenant(ctx context.Context, tenant *entities.Tenant) error
	DeleteTenant(ctx context.Context, id string) error
	GetTenantByID(ctx context.Context, id string) (*entities.Tenant, error)
	GetTenants(ctx context.Context) ([]*entities.Tenant, error)
	UpdateTenant(ctx context.Context, tenant *entities.Tenant) error
	SearchTenant(ctx context.Context, filter map[string]any) (*entities.Tenant, error)
	SearchTenants(ctx context.Context, filter map[string]any) ([]*entities.Tenant, error)
}
