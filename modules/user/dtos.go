package user

type CreateRequestDto struct {
	Name		string `json:"name" validate:"required"` 
	Email		string `json:"email" validate:"required,email"`
	Password	string `json:"password" validate:"required"`
}

type CreateResponseDto struct {
	ID			uint `json:"id"`
	Name		string `json:"name"` 
	Email		string `json:"email"`
}