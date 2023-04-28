package entities

type Permissions string

const (
	ReadPermission   Permissions = "read"
	WritePermission  Permissions = "write"
	EditPermission   Permissions = "edit"
	DeletePermission Permissions = "delete"
)
