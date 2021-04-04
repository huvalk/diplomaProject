package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func VkOAuthURL(clientID string, redirectURL string, state string) string {
	scopeTemp := "account+email"
	return fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s"+
		"&scope=%s&state=%s&display=page", clientID, redirectURL, scopeTemp, state)
}

func RetrieveUserToken(code string, clientID string, redirectURL string, clientSecret string) (*TokenStruct, error) {
	url := fmt.Sprintf("https://oauth.vk.com/access_token?grant_type=authorization_code&code=%s&"+
		"redirect_uri=%s&client_id=%s&client_secret=%s", code, redirectURL, clientID, clientSecret)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, errors.New("failed to create request to vk oauth")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("failed to make request to vk oauth")
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response from vk oauth")
	}

	token := &TokenStruct{}
	err = json.Unmarshal(bytes, token)
	if token.AccessToken == "" {
		return nil, errors.New(string(bytes))
	}

	return token, err
}

func RetrieveProfileInfo(token string) (*VKUser, error) {
	url := fmt.Sprintf("https://api.vk.com/method/%s?v=5.130&access_token=%s&fields=photo_400_orig",
		"users.get", token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("failed to create request to vk oauth")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("failed to make request to vk oauth")
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to make request to vk oauth")
	}

	user := &VKUserResponse{}
	err = json.Unmarshal(bytes, user)

	result := user.Response[0]
	return &result, err
}
