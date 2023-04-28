package entities

import (
	"github.com/hebecoding/digital-dash-commons/utils"
)

type Role struct {
	ID          utils.XID      `json:"_id"`
	TenantID    utils.XID      `json:"tenant_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Permissions []*Permissions `json:"permissions"`
}
