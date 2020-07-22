package v1

// User ...
type User struct {
	ID uint32 `json:"id"`
	// WBPassportID
	FullName           string             `json:"fullName"`
	Projects           []*Project         `json:"projects"`
	MobilePhone        string             `json:"mobilePhone"`
	UserSpecialization UserSpecialization `json:"specialization"`
	Technologies       []*Technology      `json:"technologies"`
}

// UserSpecialization ...
type UserSpecialization struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

// Users ...
type Users []*User
