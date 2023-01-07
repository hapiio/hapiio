package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"

	base "github.com/hapiio/hapiio/proto/base/v1"
	"golang.org/x/oauth2"
)

// type provider *pb.Provider

type baseProvider struct {
	provider *base.Provider
}

func (b *baseProvider) Scope() []string {
	return b.provider.GetScopes()
}

func (b *baseProvider) SetScope(scopes []string) {
	b.provider.Scopes = scopes
}

func (b *baseProvider) ClientId() string {
	return b.provider.GetClientId()
}

func (b *baseProvider) SetClientId(clientId string) {
	b.provider.ClientId = clientId
}

func (b *baseProvider) ClientSecret() string {
	return b.provider.GetClientSecret()
}

func (b *baseProvider) SetClientSecret(clientSecret string) {
	b.provider.ClientSecret = clientSecret
}

func (b *baseProvider) RedirectUrl() string {
	return b.provider.GetRedirectUrl()
}

func (b *baseProvider) SetRedirectUrl(redirectUrl string) {
	b.provider.RedirectUrl = redirectUrl
}

func (b *baseProvider) AuthUrl() string {
	return b.provider.GetAuthUrl()
}

func (b *baseProvider) SetAuthUrl(authUrl string) {
	b.provider.AuthUrl = authUrl
}

func (b *baseProvider) UserApiUrl() string {
	return b.provider.GetUserApiUrl()
}

func (b *baseProvider) SetUserApiUrl(userUrl string) {
	b.provider.UserApiUrl = userUrl
}

func (b *baseProvider) TokenUrl() string {
	return b.provider.GetTokenUrl()
}

func (b *baseProvider) SetTokenUrl(tokenUrl string) {
	b.provider.TokenUrl = tokenUrl
}

func (b *baseProvider) BuildAuthUrl(state string, opts ...oauth2.AuthCodeOption) string {
	return b.oauth2Config().AuthCodeURL(state, opts...)
}

func (b *baseProvider) FetchToken(code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return b.oauth2Config().Exchange(context.Background(), code, opts...)
}

func (b *baseProvider) Client(token *oauth2.Token) *http.Client {
	return b.oauth2Config().Client(context.Background(), token)
}

// FetchRawUserData implements Provider.FetchRawUserData interface.
func (b *baseProvider) FetchRawUserData(token *oauth2.Token) ([]byte, error) {
	req, err := http.NewRequest("GET", b.provider.GetUserApiUrl(), nil)
	if err != nil {
		return nil, err
	}

	return b.sendRawUserDataRequest(req, token)
}

// sendRawUserDataRequest sends the specified user data request and return its raw response body.
func (b *baseProvider) sendRawUserDataRequest(req *http.Request, token *oauth2.Token) ([]byte, error) {
	client := b.Client(token)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// http.Client.Get doesn't treat non 2xx responses as error
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf(
			"Failed to fetch OAuth2 user profile via %s (%d):\n%s",
			b.provider.GetUserApiUrl(),
			response.StatusCode,
			string(result),
		)
	}

	return result, nil
}

// oauth2Config constructs a oauth2.Config instance based on the provider settings.
func (b *baseProvider) oauth2Config() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  b.provider.GetRedirectUrl(),
		ClientID:     b.provider.GetClientId(),
		ClientSecret: b.provider.GetClientSecret(),
		Scopes:       b.provider.GetScopes(),
		Endpoint: oauth2.Endpoint{
			AuthURL:  b.AuthUrl(),
			TokenURL: b.TokenUrl(),
		},
	}
}
