package user_dto

type LoginUserRequest struct {
	MedicalID  *string `json:"medical_id"`
	DocumentID *string `json:"document_id"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
}

func (r *LoginUserRequest) IsDoctorLogin() bool {
	return r.MedicalID != nil && r.Password != nil
}

func (r *LoginUserRequest) IsPatientLogin() bool {
	return (r.DocumentID != nil || r.Email != nil) && r.Password != nil
}
