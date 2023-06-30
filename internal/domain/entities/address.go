package entities

type Address struct {
	ID       string `json:"_id" bson:"_id"`
	Address  string `json:"address" bson:"address"`
	Address2 string `json:"address2,omitempty" bson:"address2"`
	City     string `json:"city" bson:"city"`
	State    string `json:"state" bson:"state"`
	ZipCode  string `json:"zip_code" bson:"zip_code"`
	Country  string `json:"country" bson:"country"`
}
