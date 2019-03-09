package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bitly/go-simplejson"
)

var PublicKey = ""
var Hash = ""

func SignPayload(payload map[string]string, accessToken string) url.Values {
	v := url.Values{}
	for key, value := range payload {
		if key != "sign" {
			v.Set(key, value)
		}
	}
	v.Set("access_key", accessToken)
	v.Set("actionKey", "appkey")
	v.Set("appkey", AppKey)
	v.Set("build", "8230")
	v.Set("device", "phone")
	v.Set("mobi_app", "iphone")
	v.Set("platform", "ios")
	v.Set("ts", Timestamp())
	v.Set("type", "json")
	// v.Encode() will sort params by key
	data := v.Encode()
	v.Set("sign", Md5Sum(data+AppSecret))
	return v
}

func SignPayload2(v url.Values, accessToken string) url.Values {
	v.Set("access_key", accessToken)
	v.Set("actionKey", "appkey")
	v.Set("appkey", AppKey)
	v.Set("build", "8230")
	v.Set("device", "phone")
	v.Set("mobi_app", "iphone")
	v.Set("platform", "ios")
	v.Set("ts", Timestamp())
	v.Set("type", "json")
	// v.Encode() will sort params by key
	data := v.Encode()
	v.Set("sign", Md5Sum(data+AppSecret))
	return v
}

func GetPublicKey() (string, string, error) {
	if PublicKey != "" && Hash != "" {
		return PublicKey, Hash, nil
	}
	payload := make(map[string]string)
	resp, err := http.PostForm("https://passport.bilibili.com/api/oauth2/getKey", SignPayload(payload, ""))
	if err != nil {
		return "", "", err
	}
	j, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}
	code, err := j.Get("code").Int()
	if err != nil {
		return "", "", fmt.Errorf("cannot get result code: %s", err.Error())
	}
	if code != 0 {
		message, _ := j.Get("message").String()
		return "", "", fmt.Errorf("get public key error: %s", message)
	}
	PublicKey, err = j.Get("data").Get("key").String()
	if err != nil {
		return "", "", fmt.Errorf("cannot get public key: %s", err.Error())
	}
	Hash, err = j.Get("data").Get("hash").String()
	if err != nil {
		return "", "", fmt.Errorf("cannot get hash: %s", err.Error())
	}
	return PublicKey, Hash, nil
}

func EncryptPassword(password string) (string, error) {
	publicKey, hash, err := GetPublicKey()
	if err != nil {
		return "", fmt.Errorf("get public key failed: %s", err)
	}
	crypt, err := RsaEncrypt([]byte(publicKey), []byte(hash+password))
	if err != nil {
		return "", fmt.Errorf("rsa encrypt failed: %s", err)
	}
	passwordEncrypted := base64.StdEncoding.EncodeToString([]byte(crypt))
	return passwordEncrypted, nil
}
