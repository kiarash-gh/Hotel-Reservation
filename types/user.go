package types

type User struct{
	ID string `bson:"_id" json:"id"`
	FirstName string `bson:"fistName" json:"fistName"`
	LastName string `bson:"lastName" json:"lastName"`
} 