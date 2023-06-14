package service

import (
	"context"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/infrastructure/apperrors"
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

func (s *TenantService) UpdateTenant(ctx context.Context, id string, tenant *entities.Tenant) error {
	tenant.ID = id
	return s.Repository.UpdateTenant(ctx, tenant)
}

func (s *TenantService) DeleteTenant(ctx context.Context, id string) error {
	return s.Repository.DeleteTenant(ctx, id)
}

func (s *TenantService) GetTenants(ctx context.Context) ([]*entities.Tenant, error) {
	return s.Repository.GetTenants(ctx)
}

func (s *TenantService) GetTenantCompanies(ctx context.Context, id string) ([]*entities.TenantCompanyDetails, error) {
	var companies []*entities.TenantCompanyDetails

	tenant, err := s.Repository.GetTenantByID(ctx, id)
	if err != nil {
		s.Logger.Error(err)
		return nil, apperrors.ErrRetrievingTenantDocument
	}

	for _, company := range tenant.Companies {
		companies = append(companies, company)
	}

	if len(companies) == 0 {
		s.Logger.Info("No tenant companies found")
		return nil, apperrors.ErrNoTenantDocumentsFound
	}

	s.Logger.Infof("Found %d tenant companies", len(companies))
	return companies, nil
}
