package oauth

import "strings"

func LogoutURL() string {
	return strings.Replace(Config.Endpoint.AuthURL, "/auth", "/sessions/logout", -1)
}
