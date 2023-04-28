package repository

import (
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/internal/domain/rbac"
	"go.mongodb.org/mongo-driver/mongo"
)

type RolesRepository interface {
	SaveRole(role *rbac.Role) error
	UpdateRole(role *rbac.Role) error
	DeleteRole(roleID utils.XID) error
	FindRoleByID(roleID utils.XID) (*rbac.Role, error)
	FindAllRoles() ([]*rbac.Role, error)
}

type rolesRepositoryImpl struct {
	db *mongo.Collection
}

func NewRolesRepository(db *mongo.Collection) RolesRepository {
	return &rolesRepositoryImpl{
		db: db,
	}
}

func (r *rolesRepositoryImpl) SaveRole(role *rbac.Role) error {
	return nil
}

func (r *rolesRepositoryImpl) UpdateRole(role *rbac.Role) error {
	return nil
}

func (r *rolesRepositoryImpl) DeleteRole(roleID utils.XID) error {
	return nil
}

func (r *rolesRepositoryImpl) FindRoleByID(roleID utils.XID) (*rbac.Role, error) {
	return nil, nil
}

func (r *rolesRepositoryImpl) FindAllRoles() ([]*rbac.Role, error) {
	return nil, nil
}
