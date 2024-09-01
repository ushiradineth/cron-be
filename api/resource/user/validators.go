package user

type UserPathParams struct {
	UserID string `form:"user_id" validate:"required,uuid"`
}

type PostQueryParams struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=8,max=20,hasLowercase,hasUppercase,hasDigit,hasSpecialCharacter"`
}

type PutQueryParams struct {
	Name  string `form:"name" validate:"required"`
	Email string `form:"email" validate:"required,email"`
}
