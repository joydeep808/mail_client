package validation



type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name string `json:"name" validate:"required"`
}



type SendEmailRequest struct{
	Email string `json:"email" validate:"required,email"`
	Content string `json:"content" validate:"required"`
	UserID uint `json:"user_id" validate:"required"`
	Subject string `json:"subject" validate:"required"`
}