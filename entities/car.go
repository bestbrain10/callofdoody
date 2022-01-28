package entities

type ICars struct {
	ID      string `json:"id,omitempty" bson:"_id.omitempty"`
	Name    string `json:"name" bson:"name"`
	Model   string `json:"model" bson:"model"`
	Brand   string `json:"brand" bson:"brand"`
	Mileage string `json:"mileage" bson:"mileage"`
	IsGood  bool   `json:"isGood" bson:"isGood"`
}
