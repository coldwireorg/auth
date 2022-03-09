package hydra

type OAuth2Client struct {
	Id                      string
	Secret                  string
	GrantTypes              []string
	ResponseTypes           []string
	Scope                   []string
	Callback                string
	TokenEndpointAuthMethod string
}

/*
func CreateClient(c OAuth2Client) error {

	scope := strings.Join(c.Scope, ",")

	oAuth2Client := &h.OAuth2Client{
		ClientId:                &c.Id,
		ClientSecret:            &c.Secret,
		GrantTypes:              c.GrantTypes,
		ResponseTypes:           c.ResponseTypes,
		Scope:                   &scope,
		RedirectUris:            []string{c.Callback},
		TokenEndpointAuthMethod: &c.TokenEndpointAuthMethod,
	}

	_, _, err := Client.AdminApi.CreateOAuth2Client(context.Background()).OAuth2Client(*oAuth2Client).Execute()

	if err != nil {
		return err
	}

	oAuth2Client := *h.NewOAuth2Client() // OAuth2Client |
	resp, _, err := Client.AdminApi.CreateOAuth2Client(context.Background()).OAuth2Client(oAuth2Client).Execute()

	log.Println(resp.ClientSecret, err)

	return nil
}
*/
