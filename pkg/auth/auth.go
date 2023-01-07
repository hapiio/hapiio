package auth

import (
	"errors"
	"net/http"

	auth "github.com/hapiio/hapiio/proto/auth/v1"
	"golang.org/x/oauth2"
)

// OAuth2 client interface
type ProviderInfo interface {
	// access permissions
	Scope() []string
	// set scope
	SetScope(scopes []string)
	// provider client id
	ClientId() string
	// set provider client id
	SetClientId(clientId string)
	// client secret
	ClientSecret() string
	// set client secret
	SetClientSecret(secret string)
	//Oath Redirect URL
	RedirectUrl() string
	// set redirect url
	SetRedirectUrl(url string)
	// provider authorization service url
	AuthUrl() string
	// set provider authorization service url
	SetAuthUrl(url string)
	// Token exchange service url
	TokenUrl() string
	// set token exchange service url
	SetTokenUrl(url string)
	// user info url
	UserApiUrl() string
	// set user info url
	SetUserApiUrl(url string)
	// client
	Client(token *oauth2.Token) *http.Client
	BuildAuthUrl(state string, opts ...oauth2.AuthCodeOption) string
	FetchToken(code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	FetchRawUserData(token *oauth2.Token) ([]byte, error)
	FetchAuthUser(token *oauth2.Token) (user *auth.AuthUser, err error)
}

const (
	fb = "facebook"
)

func Provider(name string) (ProviderInfo, error) {
	switch name {
	case fb:
		return NewFacebookProvider(), nil
	default:
		return nil, errors.New("Missing Provider Name :" + name)
	}
}
