package model

type User struct {
	ID                     int    `json:"id"`
	FullName               string `json:"full_name"`
	Description            string `json:"description"`
	Email                  string `json:"email"`
	PasswordHashed         string `json:"password"`
	AvatarURl              string `json:"avatar_url"`
	HasUserTriesInstructor bool   `json:"has_user_tries_instructor"`
	PhoneNumber            string `json:"phone_number"`
}

type UserCreate struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
