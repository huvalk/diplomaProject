package oauth

type VKUserResponse struct {
	Response []VKUser `json:"response"`
}

type VKUser struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ScreenName string `json:"screen_name"`
}

type TokenStruct struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
