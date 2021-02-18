package models

type VkUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

//type OauthConfig struct {
//	ClientId     string `gorm:"column:client_id"`
//	ClientSecret string `gorm:"column:client_secret"`
//	RedirectUrl  string `gorm:"column:redirect_url"`
//}
