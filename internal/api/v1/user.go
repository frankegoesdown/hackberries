package v1

// User ...
type User struct {
	ID uint32 `json:"id"`
	//WBPassportID
	FullName string     `json:"fullName"`
	Projects []*Project `json:"projects"`
}

// UserSpecialization
type UserSpecialization struct {
}
