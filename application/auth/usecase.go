package auth

type UseCase interface {
	MakeAuthUrl(backTo string) string
	UpdateUserInfo(code string, state string, backTo string) (int, error)
	GenerateMeta(url string) (string, error)
}
