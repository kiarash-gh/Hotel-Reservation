package types

type User struct{
	ID string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"fistName" json:"fistName"`
	LastName string `bson:"lastName" json:"lastName"`
} 