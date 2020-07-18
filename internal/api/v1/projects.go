package v1

// Project ...
type Project struct {
	ID           uint32   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	OwnerID      uint32   `json:"ownerId"`
	Participants []string `json:"participants"`
	Technologies []string `json:"technologies"`
}

