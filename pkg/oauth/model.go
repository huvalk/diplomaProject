package oauth

type VKUserResponse struct {
	Response        []VKUser    `json:"response"`
}

type VKUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenStruct struct {
	AccessToken string `json:"access_token"`
}