package api

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
)

// Sign payload with access token
func (b *BilibiliApiClient) SignPayload(payload map[string]string) url.Values {
	return SignPayload(payload, b.AccessToken)
}

// Sign payload with access token
func (b *BilibiliApiClient) SignPayload2(v url.Values) url.Values {
	return SignPayload2(v, b.AccessToken)
}

// Login by username and password using OAuth2 API
// 	b.Username = "username"
// 	b.Password = "password"
// 	err := b.GetAccessToken()
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Login OK")
// 	}
func (b *BilibiliApiClient) GetAccessToken() error {
	passwordEncrypted, err := EncryptPassword(b.Password)
	if err != nil {
		return fmt.Errorf("encrypt password failed: %s", err)
	}
	payload := make(map[string]string)
	payload["seccode"] = ""
	payload["validate"] = ""
	payload["subid"] = "1"
	payload["permission"] = "ALL"
	payload["username"] = b.Username
	payload["password"] = passwordEncrypted
	payload["captcha"] = ""
	payload["challenge"] = ""
	resp, err := b.Client.PostForm("https://passport.bilibili.com/api/v2/oauth2/login", b.SignPayload(payload))
	if err != nil {
		return err
	}
	j, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return err
	}
	code, err := j.Get("code").Int()
	if err != nil {
		return fmt.Errorf("cannot get result code: %s", err)
	}
	if code != 0 {
		message, _ := j.Get("message").String()
		return fmt.Errorf("get access token failed: %s", message)
	}
	b.AccessToken, err = j.Get("data").Get("token_info").Get("access_token").String()
	if err != nil {
		return fmt.Errorf("cannot get access token: %s", err)
	}
	b.RefreshToken, err = j.Get("data").Get("token_info").Get("refresh_token").String()
	if err != nil {
		return fmt.Errorf("cannot get refresh token: %s", err)
	}
	return nil
}

// Check if the access token is valid
// 	b.AccessToken = "access token"
// 	err := b.CheckAccessToken()
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Valid access token")
// 	}
func (b *BilibiliApiClient) CheckAccessToken() error {
	payload := make(map[string]string)
	payload["access_token"] = b.AccessToken
	resp, err := b.Client.Get("https://passport.bilibili.com/api/v2/oauth2/info?" + b.SignPayload(payload).Encode())
	if err != nil {
		return err
	}
	j, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return err
	}
	code, err := j.Get("code").Int()
	if err != nil {
		return fmt.Errorf("cannot get result code: %s", err)
	}
	if code != 0 {
		message, _ := j.Get("message").String()
		return fmt.Errorf("invalid access token: %s", message)
	}
	return nil
}

// Get new access token by refresh token
// 	b.AccessToken = "invalid access token"
// 	b.RefreshToken = "refresh token"
// 	ok, message, err := b.RefreshAccessToken()
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Refresh access token OK")
// 	}
func (b *BilibiliApiClient) RefreshAccessToken() error {
	payload := make(map[string]string)
	payload["access_token"] = b.AccessToken
	payload["refresh_token"] = b.RefreshToken
	resp, err := b.Client.PostForm("https://passport.bilibili.com/api/v2/oauth2/refreshToken", b.SignPayload(payload))
	if err != nil {
		return err
	}
	j, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return err
	}
	code, err := j.Get("code").Int()
	if err != nil {
		return fmt.Errorf("cannot get result code: %s", err)
	}
	if code != 0 {
		message, _ := j.Get("message").String()
		return errors.New("refresh access token failed: " + message)
	}
	return nil
}

// Get cookies after OAuth2 login. Must be called after LoginByUsernamePassword, LoginByAccessToken or LoginByRefreshToken.
func (b *BilibiliApiClient) GetCookies() error {
	payload := make(map[string]string)
	_, err := b.Client.Get("https://passport.bilibili.com/api/login/sso?" + b.SignPayload(payload).Encode())
	if err != nil {
		return err
	}
	return nil
}

// Check if the cookies is valid
func (b *BilibiliApiClient) CheckCookies() error {
	v := url.Values{}
	v.Set("ts", strconv.FormatInt(time.Now().Unix(), 10))
	resp, err := b.Client.Get("https://api.live.bilibili.com/User/getUserInfo?" + v.Encode())
	if err != nil {
		return err
	}
	j, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return err
	}
	code, err := j.Get("code").String()
	if err != nil {
		message, _ := j.Get("message").String()
		return errors.New("check cookie failed: " + message)
	}
	if code != "REPONSE_OK" {
		message, _ := j.Get("message").String()
		return errors.New("cookie expired: " + message)
	}
	return nil
}

func (b *BilibiliApiClient) LoginByUsernamePassword(username string, password string) error {
	if username == "" {
		return errors.New("empty username")
	}
	if password == "" {
		return errors.New("empty password")
	}
	b.Username = username
	b.Password = password
	return b.GetAccessToken()
}

func (b *BilibiliApiClient) LoginByAccessToken(accessToken string) error {
	if accessToken == "" {
		return errors.New("empty access token")
	}
	b.AccessToken = accessToken
	err := b.CheckAccessToken()
	if err != nil {
		return err
	}
	return nil
}

func (b *BilibiliApiClient) LoginByRefreshToken(accessToken string, refreshToken string) error {
	if accessToken == "" {
		return errors.New("empty access token")
	}
	if refreshToken == "" {
		return errors.New("empty refresh token")
	}
	b.AccessToken = accessToken
	b.RefreshToken = refreshToken
	return b.RefreshAccessToken()
}

// Login by several step
//
// 1. Set the cookie.
//
// 2. Login by access token
//
// 3. If access token invalid, login by refresh token
//
// 4. If refresh token invalid, login by username and password
//
// 5. Check current cookies. If cookies are invalid, try get new cookies by access token.
func (b *BilibiliApiClient) Login(username string, password string, accessToken string, refreshToken string, jsonCookie []byte) error {
	var err error
	b.Username = username
	b.Password = password
	b.AccessToken = accessToken
	b.RefreshToken = refreshToken
	_ = b.LoadCookies(jsonCookie)
	err = b.LoginByAccessToken(accessToken)
	if err != nil {
		err = b.LoginByRefreshToken(accessToken, refreshToken)
		if err != nil {
			err = b.LoginByUsernamePassword(username, password)
			if err != nil {
				return err
			}
		}
	}
	err = b.CheckCookies()
	if err != nil {
		err = b.GetCookies()
		if err != nil {
			return err
		}
	}
	return nil
}
