package validation

type SignupForm struct {
	Email    string `validate:"required,email" binding:"required" json:"email"`
	Password string `validate:"required,min=8" binding:"required" json:"password"`
	Name     string `validate:"required,max=100" binding:"required" json:"name"`
}

type LoginForm struct {
	Email    string `validate:"required" binding:"required" json:"email"`
	Password string `validate:"required" binding:"required" json:"password"`
}
