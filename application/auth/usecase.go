package auth

type UseCase interface {
	MakeAuthUrl() string
	UpdateUserInfo(code string, state string) (int, error)
}
