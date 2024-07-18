package user_dto

type CreateUserRequest struct {
	FullName   string `json:"full_name" validate:"required,min=5,max=255"`
	Email      string `json:"email" validate:"required,email,min=5,max=255"`
	Password   string `json:"password" validate:"required,min=8,max=30"`
	DocumentID string `json:"document_id" validate:"required,min=11,max=14,cpfcnpj"`
	Phone      string `json:"phone" validate:"required,min=5,max=255"`
	Role       string `json:"role" validate:"required,oneof=doctor patient"`

	DoctorMedicalID string  `json:"medical_id"`
	DoctorSpecialty string  `json:"specialty"`
	DoctorPrice     float64 `json:"price"`
}
