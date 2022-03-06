package hydra

import (
	"context"
)

func Logout(challenge string) (string, error) {
	resp, _, err := Client.AdminApi.AcceptLogoutRequest(context.Background()).LogoutChallenge(challenge).Execute()
	if err != nil {
		return "/sign-in", err
	}

	return resp.RedirectTo, nil
}
