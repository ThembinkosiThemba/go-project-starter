package user

type USER struct {
	ID      int    `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name" validate:"required"`
	Surname string `json:"surname" bson:"surname" validate:"required"`
	Email   string `json:"email" bson:"email" validate:"required"`
}