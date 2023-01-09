package auth

import (
	"encoding/json"

	auth "github.com/hapiio/hapiio/proto/auth/v1"
	base "github.com/hapiio/hapiio/proto/base/v1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var _ ProviderInfo = (*Facebook)(nil)

// NameFacebook is the unique name of the Facebook provider.
const NameFacebook string = "facebook"

// Facebook allows authentication via Facebook OAuth2.
type Facebook struct {
	*baseProvider
}

// NewFacebookProvider creates new Facebook provider instance with some defaults.
func NewFacebookProvider() *Facebook {

	provider := &baseProvider{provider: &base.Provider{
		Scopes:     []string{"email"},
		AuthUrl:    facebook.Endpoint.AuthURL,
		TokenUrl:   facebook.Endpoint.TokenURL,
		UserApiUrl: "https://graph.facebook.com/me?fields=name,email,picture.type(large)",
	}}

	return &Facebook{provider}
}

// FetchAuthUser returns an AuthUser instance based on the Facebook's user api.
//
// API reference: https://developers.facebook.com/docs/graph-api/reference/user/
func (p *Facebook) FetchAuthUser(token *oauth2.Token) (*auth.AuthUser, error) {
	data, err := p.FetchRawUserData(token)
	if err != nil {
		return nil, err
	}

	rawUser := map[string]string{}
	if err := json.Unmarshal(data, &rawUser); err != nil {
		return nil, err
	}

	extracted := struct {
		Id      string
		Name    string
		Email   string
		Picture struct {
			Data struct{ Url string }
		}
	}{}
	if err := json.Unmarshal(data, &extracted); err != nil {
		return nil, err
	}

	user := &auth.AuthUser{
		Id:          extracted.Id,
		Name:        extracted.Name,
		Email:       extracted.Email,
		AvatarUrl:   extracted.Picture.Data.Url,
		RawUser:     rawUser,
		AccessToken: token.AccessToken,
	}

	return user, nil
}
