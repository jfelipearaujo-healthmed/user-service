package user_dto

type UpdateUserRequest struct {
	FullName   *string `json:"full_name" validate:"omitempty,min=5,max=255"`
	Email      *string `json:"email" validate:"omitempty,email,min=5,max=255"`
	Password   *string `json:"password" validate:"omitempty,min=8,max=30"`
	DocumentID *string `json:"document_id" validate:"omitempty,min=11,max=14,cpfcnpj"`
	Phone      *string `json:"phone" validate:"omitempty,min=5,max=255"`

	DoctorMedicalID *string  `json:"medical_id" validate:"omitempty,required,min=5,max=255"`
	DoctorSpecialty *string  `json:"specialty" validate:"omitempty,required,min=5,max=255"`
	DoctorPrice     *float64 `json:"price" validate:"omitempty,required,min=1"`
}
