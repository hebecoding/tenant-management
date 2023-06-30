package entities

type Role struct {
	ID          string       `json:"_id" bson:"_id"`
	Name        string       `json:"name" bson:"name"`
	Description string       `json:"description" bson:"description"`
	Permissions []Permission `json:"permissions" bson:"permissions"`
}
