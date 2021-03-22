package models

//db model
type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	WorkPlace   string `json:"workPlace"`
	Description string `json:"description"`
	Bio         string `json:"bio"`
	Avatar      string `json:"avatar"`
}

//type OauthConfig struct {
//	ClientId     string `gorm:"column:client_id"`
//	ClientSecret string `gorm:"column:client_secret"`
//	RedirectUrl  string `gorm:"column:redirect_url"`
//}

//easyjson:json
type UserArr []User

type FeedUser struct {
	Id          int      `json:"id"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Email       string   `json:"email"`
	WorkPlace   string   `json:"workPlace"`
	Description string   `json:"description"`
	Bio         string   `json:"bio"`
	Avatar      string   `json:"avatar"`
	Tm          Team     `json:"team"`
	JobName     string   `json:"jobName"`
	Skills      []Skills `json:"skills"`
}

//easyjson:json
type FeedUserArr []FeedUser

func (fu *FeedUser) Convert(usr User) {
	fu.Id = usr.Id
	fu.FirstName = usr.FirstName
	fu.LastName = usr.LastName
	fu.Email = usr.Email
	fu.Description = usr.Description
	fu.WorkPlace = usr.WorkPlace
	fu.Bio = usr.Bio
	fu.Avatar = usr.Avatar
}
