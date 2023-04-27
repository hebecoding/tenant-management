package rbac

import "github.com/hebecoding/digital-dash-commons/utils"

type PermissionService interface {
	CreatePermission(permission *Permission) error
	UpdatePermission(permission *Permission) error
	DeletePermission(id utils.XID) error
	GetPermission(id utils.XID) (*Permission, error)
	GetPermissions() ([]*Permission, error)
}

type permissionServiceImp struct{}

func NewPermissionService() PermissionService {
	return &permissionServiceImp{}
}

func (p *permissionServiceImp) CreatePermission(permission *Permission) error {
	return nil
}

func (p *permissionServiceImp) UpdatePermission(permission *Permission) error {
	return nil
}

func (p *permissionServiceImp) DeletePermission(id utils.XID) error {
	return nil
}

func (p *permissionServiceImp) GetPermission(id utils.XID) (*Permission, error) {
	return nil, nil
}

func (p *permissionServiceImp) GetPermissions() ([]*Permission, error) {
	return nil, nil
}
