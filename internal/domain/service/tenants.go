package service

import (
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/internal/domain/repository"
)

type TenantService struct {
	Repository repository.TenantRepository
	Logger     utils.LoggerInterface
}

func NewTenantService(
	repository repository.TenantRepository,
	logger utils.LoggerInterface,
) *TenantService {
	return &TenantService{
		Repository: repository,
		Logger:     logger,
	}
}
