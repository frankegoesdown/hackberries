package v1

// Technology ...
type Technology struct {
	ID   uint32 `json:"id" example:"1"`
	Name string `json:"name" example:"golang"`
	Kind string `json:"kind" example:"back"`
}
