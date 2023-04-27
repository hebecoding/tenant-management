package rbac

import "github.com/hebecoding/digital-dash-commons/utils"

type RoleService interface {
	CreateRole(role *Role) error
	UpdateRole(role *Role) error
	DeleteRole(id utils.XID) error
	GetRole(id utils.XID) (*Role, error)
	GetRoles() ([]*Role, error)
}

type roleServiceImp struct{}

func NewRoleService() RoleService {
	return &roleServiceImp{}
}

func (r *roleServiceImp) CreateRole(role *Role) error {
	return nil
}

func (r *roleServiceImp) UpdateRole(role *Role) error {
	return nil
}

func (r *roleServiceImp) DeleteRole(id utils.XID) error {
	return nil
}

func (r *roleServiceImp) GetRole(id utils.XID) (*Role, error) {
	return nil, nil
}

func (r *roleServiceImp) GetRoles() ([]*Role, error) {
	return nil, nil
}
