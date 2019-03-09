package api

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/ganlvtech/go-exportable-cookiejar"
	"golang.org/x/net/publicsuffix"
)

type BilibiliApiClient struct {
	Client        http.Client
	Username      string
	Password      string
	AccessToken   string
	RefreshToken  string
	biliJct       string
	DanmakuConfig struct{
		Color  int
		Mode   int
		Length int
	}
}

func NewBilibiliApiClient(debug bool) *BilibiliApiClient {
	b := new(BilibiliApiClient)
	b.Client = http.Client{}
	b.Client.Jar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	if debug {
		proxyStr := "http://localhost:8888"
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			panic(err)
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		b.Client.Transport = transport
	}

	return b
}

func (b *BilibiliApiClient) BiliJct() (string, error) {
	if b.biliJct != "" {
		return b.biliJct, nil
	}
	u, _ := url.Parse("https://api.live.bilibili.com")
	cookies := b.Client.Jar.Cookies(u)
	for _, cookie := range cookies {
		if cookie.Name == "bili_jct" {
			b.biliJct = cookie.Value
			return b.biliJct, nil
		}
	}
	return "", errors.New("cannot find bili_jct in cookies")
}

func (b *BilibiliApiClient) SaveCookie() ([]byte, error) {
	j, ok := b.Client.Jar.(*cookiejar.Jar)
	if !ok {
		return []byte{}, errors.New("cookie jar type assertion failed")
	}
	return j.JsonSerialize()
}

func (b *BilibiliApiClient) LoadCookies(data []byte) error {
	j, ok := b.Client.Jar.(*cookiejar.Jar)
	if !ok {
		return errors.New("cookie jar type assertion failed")
	}
	return j.JsonDeserialize(data)
}
