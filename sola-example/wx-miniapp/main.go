package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"

	"github.com/ddosakura/sola/v2/middleware/router"

	"github.com/ddosakura/sola/v2"
)

// main

func main() {
	app := sola.New()
	app.LoadConfig()
	appID := viper.GetString("wx.appid")
	secret := viper.GetString("wx.secret")
	code2sessionTpl = fmt.Sprintf(code2sessionTpl2, appID, secret, "%s")
	{
		r := router.New(nil)
		r.Bind("GET /onLogin", onLogin)
		app.Use(r.Routes())
	}
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}

// meta

const (
	code2sessionTpl2 = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

var (
	code2sessionTpl string
)

// utils

type sessBody struct {
	ErrCode int
	ErrMsg  string

	OpenID     string
	SessionKey string `json:"session_key"`
	// UnionID    string
}

func code2session(code string) (*sessBody, error) {
	url := fmt.Sprintf(code2sessionTpl, code)
	var e error
	var res *http.Response
	if res, e = http.Get(url); e != nil {
		return nil, e
	}
	defer res.Body.Close()
	var bs []byte
	if bs, e = ioutil.ReadAll(res.Body); e != nil {
		return nil, e
	}

	var sb sessBody
	if e = json.Unmarshal(bs, &sb); e != nil {
		return nil, e
	}
	return &sb, nil
}

// model

type resBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// controller

func onLogin(c sola.Context) error {
	r := c.Request()
	code := r.URL.Query().Get("code")
	var e error
	var data *sessBody
	if data, e = code2session(code); e != nil {
		return c.JSON(http.StatusOK, &resBody{
			Code: -1,
			Msg:  "code2session error",
		})
	}
	if data.ErrCode != 0 {
		return c.JSON(http.StatusOK, &resBody{
			Code: -1,
			Msg:  data.ErrMsg,
		})
	}

	// login/register
	u := users[data.OpenID]
	if u == nil {
		u = &user{
			OpenID:     data.OpenID,
			SessionKey: data.SessionKey,
			LoginTimes: 1,
		}
		users[data.OpenID] = u
	} else {
		u.LoginTimes++
	}
	return c.JSON(http.StatusOK, &resBody{
		Code: 0,
		Msg:  "success",
		Data: u,
	})
}

// mock

type user struct {
	OpenID     string `json:"-"`
	SessionKey string `json:"-"`
	LoginTimes int    `json:"times"`
}

var (
	users = make(map[string]*user)
)
