package user

// USER represents a user entity in the system.
// It contains basic information about a user such as their ID, name, surname, and email.
type USER struct {
    // ID is the unique identifier for the user.
    // It's represented as a string in JSON and BSON.
    ID      string `json:"id" bson:"id"`

    // Name is the user's first name.
    // It's a required field for both JSON and BSON.
    Name    string `json:"name" bson:"name" validate:"required"`

    // Surname is the user's last name.
    // It's a required field for both JSON and BSON.
    Surname string `json:"surname" bson:"surname" validate:"required"`

    // Email is the user's email address.
    // It's a required field for both JSON and BSON.
    Email   string `json:"email" bson:"email" validate:"required"`
}