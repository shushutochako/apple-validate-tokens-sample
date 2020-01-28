package signinapple

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"shushutochako/pkg/key"
	"strings"
	"time"

	"shushutochako/pkg/config"
)

type Client struct {
}

type ValidateTokensParams struct {
	AuthorizationCode string `json:"authorization_code"`
}

type ValidateTokensResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}

type ValidateTokensErrorResponse struct {
	Error string `json:"error"`
}

const TimeoutDuration = 15

func NewClient() *Client {
	return &Client{}
}

/// TODO
func (_client Client) ValidateTokens(params ValidateTokensParams) (ValidateTokensResponse, error) {
	config := config.NewConfig()
	keys := key.NewKey()
	cs, error := NewClientSecret(config, keys)
	if error != nil {
		return ValidateTokensResponse{}, error
	}

	form := url.Values{}
	form.Add("client_id", config.ClientID)
	form.Add("client_secret", cs.value)
	form.Add("code", params.AuthorizationCode)
	form.Add("grant_type", "authorization_code")
	body := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", "https://appleid.apple.com/auth/token", body)
	if err != nil {
		return ValidateTokensResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{Timeout: time.Duration(TimeoutDuration) * time.Second}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return ValidateTokensResponse{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var tBody bytes.Buffer
		tee := io.TeeReader(resp.Body, &tBody)
		vter := ValidateTokensErrorResponse{}
		json.NewDecoder(tee).Decode(&vter)
		fmt.Println("error")
		fmt.Println(vter.Error)
		return ValidateTokensResponse{}, err
	} else {
		var tBody bytes.Buffer
		tee := io.TeeReader(resp.Body, &tBody)
		vts := ValidateTokensResponse{}
		json.NewDecoder(tee).Decode(&vts)
		fmt.Println("tokenType")
		fmt.Println(vts.TokenType)
		return ValidateTokensResponse{}, nil
	}

	// 検証リクエスト
	/*
		form := url.Values{}
		form.Add("client_id", config.Apple.ClientID)
		form.Add("client_secret", cs.value)
		form.Add("code", params.AuthorizationCode)
		form.Add("grant_type", "authorization_code")

		body := strings.NewReader(form.Encode())
		req, err := http.NewRequest("POST", "https://appleid.apple.com/auth/token", body)
		if err != nil {
			return ValidateTokensResponse{}, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{Timeout: time.Duration(TimeoutDuration) * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return ValidateTokensResponse{}, err
		}

		defer resp.Body.Close()

		var tBody bytes.Buffer
		tee := io.TeeReader(resp.Body, &tBody)
		vts := ValidateTokensResponse{}
		json.NewDecoder(tee).Decode(&vts)

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			he := httpClient.HTTPError{StatusCode: resp.StatusCode, Body: tBody.String()}
			logger.Error(he.Body)
			return vts, &he
		}
		fmt.Println(vts.TokenType)
		return vts, nil
	*/
}
