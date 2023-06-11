package service

import (
	"context"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	"github.com/hebecoding/tenant-management/internal/domain/repository"
)

type TenantService struct {
	Repository repository.TenantRepository
	Logger     utils.LoggerInterface
}

func NewTenantService(
	logger utils.LoggerInterface,
	repository repository.TenantRepository,
) *TenantService {
	return &TenantService{
		Repository: repository,
		Logger:     logger,
	}
}

func (s *TenantService) CreateTenant(ctx context.Context, tenant *entities.Tenant) error {
	return s.Repository.CreateTenant(ctx, tenant)
}

func (s *TenantService) GetTenantByID(ctx context.Context, id string) (*entities.Tenant, error) {
	return s.Repository.GetTenantByID(ctx, id)
}
