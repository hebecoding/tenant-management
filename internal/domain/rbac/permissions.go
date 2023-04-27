package rbac

import "github.com/hebecoding/digital-dash-commons/utils"

type Permission struct {
	ID          utils.XID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
