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

func (s *TenantService) GetTenantCompaniesSubscriptions(ctx context.Context, id string) (
	[]*entities.TenantSubscriptionDetails, error,
) {
	var subscriptions []*entities.TenantSubscriptionDetails

	tenant, err := s.Repository.GetTenantByID(ctx, id)
	if err != nil {
		s.Logger.Error(err)
		return nil, apperrors.ErrRetrievingTenantDocument
	}

	for _, company := range tenant.Companies {
		for _, subscription := range company.Subscriptions {
			subscriptions = append(subscriptions, subscription)
		}
	}

	if len(subscriptions) == 0 {
		s.Logger.Info("no tenant subscriptions found")
		return nil, apperrors.ErrNoTenantDocumentsFound
	}

	s.Logger.Infof("found %d tenant subscriptions", len(subscriptions))
	return subscriptions, nil
}

func (s *TenantService) UpdateTenantSubscription(
	ctx context.Context, tenantID string, subscription *entities.TenantSubscriptionDetails,
) error {
	var tenant = &entities.Tenant{}

	if subscription == nil {
		s.Logger.Info("subscription is nil")
		return apperrors.ErrInvalidTenantSubscription
	}

	// Get the t by ID
	s.Logger.Infof("getting tenant by ID: %s", tenantID)
	t, err := s.GetTenantByID(ctx, tenantID)
	if err != nil {
		return err
	}

	// Update the subscription
	tenant.ID = tenantID

	found := false
outer:
	for _, company := range t.Companies {
		for _, sub := range company.Subscriptions {
			if sub.ID == subscription.ID {
				tenant.Companies = append(
					tenant.Companies, &entities.TenantCompanyDetails{
						Subscriptions: []*entities.TenantSubscriptionDetails{subscription},
					},
				)
				found = true
				break outer
			}
		}
	}

	if !found {
		s.Logger.Infof("subscription with ID %s not found", subscription.ID)
		return apperrors.ErrNoTenantDocumentsFound
	}

	return s.Repository.UpdateTenant(ctx, tenant)
}
