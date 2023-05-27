package entities

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

type Permission string

const (
	ReadPermission   Permission = "read"
	WritePermission  Permission = "write"
	EditPermission   Permission = "edit"
	DeletePermission Permission = "delete"
)

func (p *Permission) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = strings.ToLower(str)
	str = str[1 : len(str)-1] // Remove the quotes from the JSON string

	switch str {
	case string(ReadPermission):
		*p = ReadPermission
	case string(WritePermission):
		*p = WritePermission
	case string(EditPermission):
		*p = EditPermission
	case string(DeletePermission):
		*p = DeletePermission
	default:
		return errors.New("Invalid permission")
	}
	return nil
}

func (p Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}
