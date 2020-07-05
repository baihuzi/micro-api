package security

type AccessTokenGeneratorInterface interface {
	Create(arg ...string) (token string)
}
