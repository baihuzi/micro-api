package security

import "net/http"

type FirewallInterface interface {
	GetCredentialsFromRequest(r *http.Request) (jwt string)
	IfPattern(path string) bool
	Authenticate(jwtStr string) error
}
