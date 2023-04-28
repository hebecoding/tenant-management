package service

import (
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
)

type RoleService interface {
	CreateRole(role *entities.Role) error
	CreateCustomRole(role *entities.Role) error
	UpdateRole(role *entities.Role) error
	DeleteRole(id utils.XID) error
	GetRole(id utils.XID) (*entities.Role, error)
	GetRoles() ([]*entities.Role, error)
}

type roleServiceImp struct{}

func NewRoleService() RoleService {
	return &roleServiceImp{}
}

func (r *roleServiceImp) CreateRole(role *entities.Role) error {
	return nil
}

func (r *roleServiceImp) CreateCustomRole(role *entities.Role) error {
	return nil
}

func (r *roleServiceImp) UpdateRole(role *entities.Role) error {
	return nil
}

func (r *roleServiceImp) DeleteRole(id utils.XID) error {
	return nil
}

func (r *roleServiceImp) GetRole(id utils.XID) (*entities.Role, error) {
	return nil, nil
}

func (r *roleServiceImp) GetRoles() ([]*entities.Role, error) {
	return nil, nil
}
