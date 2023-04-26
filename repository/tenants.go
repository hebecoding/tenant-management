package repository

type TenantRepository interface {
}
type TenantRepositoryImpl struct {
}

func NewTenantRepository() *TenantRepositoryImpl {
	return &TenantRepositoryImpl{}
}
