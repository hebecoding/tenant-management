package repository

import (
	"context"

	"github.com/hebecoding/tenant-management/internal/domain/entities"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *entities.Tenant) error
	GetTenantByID(ctx context.Context, id string) (*entities.Tenant, error)
	UpdateTenant(ctx context.Context, tenant *entities.Tenant) (*entities.Tenant, error)
	DeleteTenant(ctx context.Context, id string) error
	ListTenants(ctx context.Context, offset, limit int64) ([]*entities.Tenant, error)
	GetTenantExtraInfo(ctx context.Context, id string, infoType string) (interface{}, error)
	UpdateTenantExtraInfo(ctx context.Context, id string, infoType string, info interface{}) error
	UpdateTenantFields(ctx context.Context, id string, fieldsToUpdate map[string]interface{}) (*entities.Tenant, error)
}
