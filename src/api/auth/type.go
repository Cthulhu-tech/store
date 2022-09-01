package auth

type VK_DATA struct {
	AT      string `json:"access_token"`
	EMAIL   string `json:"email"`
	EXPIRES int    `json:"expires_in"`
	USER_ID int    `json:"user_id"`
}

type UserRegistration struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegistrationFunc struct {
	Value int `json:"sp_registation"`
}

type UserRegistrationConfirm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Secret   int    `json:"secret"`
}

type ConfirmVkFunc struct {
	Value int `json:"sp_cofnfirm_vk"`
}

type UserConfirm struct {
	Code   string `json:"code"`
	Secret int    `json:"secret"`
}

type ConfirmFunc struct {
	Value int `json:"sp_confirm"`
}

type UserLogin struct {
	Login    string `json:"login"`
	Mail     string `json:"email"`
	Password string `json:"password"`
}

type UserAllData struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  bool   `json:"confirme"`
}

type MessageToken struct {
	Message string
	Token   string
}
