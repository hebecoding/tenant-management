package entities

import (
	"github.com/hebecoding/digital-dash-commons/utils"
)

type Role struct {
	ID          utils.XID     `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Permissions []*Permission `json:"permissions" bson:"permissions"`
}
