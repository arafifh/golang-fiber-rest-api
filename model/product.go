package model

type Product struct {
	ID    string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string  `json:"name" bson:"name"`
	Price int 	  `json:"price" bson:"price"`
}