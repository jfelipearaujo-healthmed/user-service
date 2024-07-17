package health

type HealthStatus struct {
	Status string `json:"status"`
	Err    string `json:"err,omitempty"`
}

func New(err error) *HealthStatus {
	status := "ok"
	if err != nil {
		status = "error"
	}

	return &HealthStatus{
		Status: status,
		Err:    err.Error(),
	}
}
