package auth

type VK_DATA struct {
	AT      string `json:"access_token"`
	EMAIL   string `json:"email"`
	EXPIRES int    `json:"expires_in"`
	USER_ID int    `json:"user_id"`
}

type Error struct {
	message string
}
