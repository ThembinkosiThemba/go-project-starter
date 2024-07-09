package dto

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Delete struct {
	Email string `json:"email" validate:"required"`
}

type Profile struct {
	Email string `json:"email" validate:"required"`
}

type ResetPassword struct {
	Email       string `json:"email" validate:"required"`
	NewPassword string `json:"password" validate:"required"`
}
