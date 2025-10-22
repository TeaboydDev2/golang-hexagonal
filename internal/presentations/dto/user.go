package dto

type CraeteUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}
