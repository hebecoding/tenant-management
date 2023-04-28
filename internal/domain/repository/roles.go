package repository

import (
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type RolesRepository interface {
	SaveRole(role *entities.Role) error
	UpdateRole(role *entities.Role) error
	DeleteRole(roleID utils.XID) error
	FindRoleByID(roleID utils.XID) (*entities.Role, error)
	FindAllRoles() ([]*entities.Role, error)
}

type rolesRepositoryImpl struct {
	db *mongo.Collection
}

func NewRolesRepository(db *mongo.Collection) RolesRepository {
	return &rolesRepositoryImpl{
		db: db,
	}
}

func (r *rolesRepositoryImpl) SaveRole(role *entities.Role) error {
	return nil
}

func (r *rolesRepositoryImpl) UpdateRole(role *entities.Role) error {
	return nil
}

func (r *rolesRepositoryImpl) DeleteRole(roleID utils.XID) error {
	return nil
}

func (r *rolesRepositoryImpl) FindRoleByID(roleID utils.XID) (*entities.Role, error) {
	return nil, nil
}

func (r *rolesRepositoryImpl) FindAllRoles() ([]*entities.Role, error) {
	return nil, nil
}
