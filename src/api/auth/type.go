package auth

type VK_DATA struct {
	AT      string `json:"access_token"`
	EMAIL   string `json:"email"`
	EXPIRES int    `json:"expires_in"`
	USER_ID int    `json:"user_id"`
}

type Message struct {
	message string
}

type UserRegistration struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegistrationFunc struct {
	Value int `json:"sp_registation"`
}
