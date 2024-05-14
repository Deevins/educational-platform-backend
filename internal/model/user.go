package model

type User struct {
	ID          int32  `json:"id"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Avatar      []byte `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
	Linkedin    string `json:"linkedin"`
	VK          string `json:"vk"`
	Telegram    string `json:"telegram"`
	Github      string `json:"github"`
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
	Linkedin    string `json:"linkedin"`
	VK          string `json:"vk"`
	Github      string `json:"github"`
}

type UserUpdateAvatar struct {
	UserID int32  `json:"user_id"`
	Avatar []byte `json:"avatar"`
}

type UserUpdateTeachingExperience struct {
	UserID                int32  `json:"user_id"`
	HasVideoKnowledge     string `json:"has_video_knowledge"`
	HasPreviousExperience string `json:"has_previous_experience"`
	CurrentAudienceCount  string `json:"current_audience_count"`
}
