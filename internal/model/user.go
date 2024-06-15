package model

type User struct {
	ID          int32  `json:"id"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	AvatarUrl   string `json:"avatar_url"`
	PhoneNumber string `json:"phone_number"`
}

type UserCreate struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserUpdate struct {
	UserID      int32  `json:"user_id"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type UserUpdateAvatar struct {
	UserID    int32  `json:"user_id"`
	AvatarUrl []byte `json:"avatar_url"`
}

type UserUpdateTeachingExperience struct {
	UserID             int32  `json:"user_id"`
	VideoKnowledge     string `json:"video_knowledge"`
	PreviousExperience string `json:"previous_experience"`
	CurrentAudience    string `json:"current_audience"`
}

type UserIDWithResourceLink struct {
	UserID int32  `json:"user_id"`
	Link   string `json:"link"`
}
