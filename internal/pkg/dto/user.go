package dto

type User struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber int64  `json:"phone_number"`
	UserType    string `json:"user_type"`
}

type CreateUserRequest struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber int64  `json:"phone_number"`
	UserType    string `json:"user_type"`
}
