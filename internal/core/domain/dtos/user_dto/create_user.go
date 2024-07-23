package user_dto

type CreateUserRequest struct {
	FullName   string `json:"full_name" validate:"required,min=5,max=255"`
	Email      string `json:"email" validate:"required,email,min=5,max=255"`
	Password   string `json:"password" validate:"required,min=8,max=30"`
	DocumentID string `json:"document_id" validate:"required,min=11,max=14,cpfcnpj"`
	Phone      string `json:"phone" validate:"required,min=5,max=255"`
	Role       string `json:"role" validate:"required,oneof=doctor patient"`

	Doctor  *CreateDoctorRequest  `json:"doctor" validate:"omitempty"`
	Address *CreateAddressRequest `json:"address" validate:"omitempty"`
}

type CreateDoctorRequest struct {
	MedicalID string  `json:"medical_id" validate:"omitempty,required,min=5,max=255"`
	Specialty string  `json:"specialty" validate:"omitempty,required,min=5,max=255"`
	Price     float64 `json:"price" validate:"omitempty,required,min=1"`
}

type CreateAddressRequest struct {
	Street       string `json:"street" validate:"required,min=1,max=255"`
	Number       string `json:"number" validate:"required,min=1,max=255"`
	Neighborhood string `json:"neighborhood" validate:"required,min=1,max=255"`
	City         string `json:"city" validate:"required,min=1,max=255"`
	State        string `json:"state" validate:"required,min=1,max=255"`
	Zip          string `json:"zip" validate:"required,min=1,max=255"`
}
