package repository

import (
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
)

type RolesRepository interface {
	SaveRole(role *entities.Role) error
	UpdateRole(role *entities.Role) error
	DeleteRole(roleID utils.XID) error
	FindRoleByID(roleID utils.XID) (*entities.Role, error)
	FindAllRoles() ([]*entities.Role, error)
}
