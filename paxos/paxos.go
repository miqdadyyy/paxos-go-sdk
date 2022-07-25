package paxos

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/miqdadyyy/paxos-go-sdk/constant"
	"github.com/miqdadyyy/paxos-go-sdk/paxos/oauth"
	"net/http"
	"strings"
	time "time"
)

type PaxosClient struct {
	ClientID        string     `json:"client_id"`
	ClientSecret    string     `json:"client_secret"`
	IsAuthenticated bool       `json:"is_authenticated"`
	BaseURL         string     `json:"base_url"`
	OauthBaseURL    string     `json:"oauth_base_url"`
	OauthData       PaxosOauth `json:"oauth"`
}

type PaxosOauth struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	Scope       string    `json:"scope"`
	ExpireTime  time.Time `json:"expire_time"`
}

func New(clientID, clientSecret string, sandbox bool) PaxosClient {
	client := PaxosClient{
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		IsAuthenticated: false,
	}

	if sandbox {
		client.BaseURL = constant.PaxosSandboxBaseURL
		client.OauthBaseURL = constant.PaxosSandboxOauthBaseURL
	} else {
		client.BaseURL = constant.PaxosBaseURL
		client.OauthBaseURL = constant.PaxosOauthBaseURL
	}

	return client
}

func (c *PaxosClient) GenerateClientRequest() *resty.Request {
	client := resty.New()
	if c.IsAuthenticated && c.OauthData.ExpireTime.After(time.Now()) {
		client = client.SetAuthScheme(c.OauthData.TokenType).
			SetAuthToken(c.OauthData.AccessToken)
	}

	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if response.StatusCode() != http.StatusOK {
			return errors.New(response.String())
		}

		// TODO : RE Auth when response is 401

		return nil
	})

	return client.R()
}

func (c *PaxosClient) OAuth(scopes ...string) error {
	client := c.GenerateClientRequest()
	url := fmt.Sprintf("%s%s", c.OauthBaseURL, "oauth2/token")
	resp, err := client.
		SetFormData(map[string]string{
			"grant_type":    constant.PaxosGrantTypeClientCredentials,
			"client_id":     c.ClientID,
			"client_secret": c.ClientSecret,
			"scope":         strings.Join(scopes, " "),
		}).
		Post(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(string(resp.Body()))
	}

	var oauthResp oauth.OAuthResponse
	if err := json.Unmarshal(resp.Body(), &oauthResp); err != nil {
		return err
	}

	c.OauthData = PaxosOauth{
		AccessToken: oauthResp.AccessToken,
		TokenType:   oauthResp.TokenType,
		Scope:       oauthResp.Scope,
		ExpireTime:  time.Now().Add(time.Second * time.Duration(oauthResp.ExpiresIn)),
	}
	c.IsAuthenticated = true
	return nil
}
