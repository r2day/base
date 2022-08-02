package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/r2day/base/conf"
	logger "github.com/r2day/base/log"
)

const (
	argsTpl = "%s?appid=%s&secret=%s&js_code=%s&grant_type=%s"
)

// Code2Session 微信的session获取
func Code2Session(code string) ([]byte, error) {
	logger.Logger.WithField("code", code)
	code2SessionUrl := conf.ConfInstance.Wx.Code2SessionUrl
	appId := conf.ConfInstance.Wx.AppId
	secret := conf.ConfInstance.Wx.Secret
	grantType := conf.ConfInstance.Wx.GrantType

	// 渲染参数
	args := fmt.Sprintf(argsTpl,
		code2SessionUrl,
		appId,
		secret,
		code,
		grantType,
	)

	resp, err := http.Get(args)
	if err != nil {
		logger.Logger.WithError(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.WithError(err)
		return nil, err
	}

	logger.Logger.WithField("body", string(body)).
		Println("before exit code 2session")
	return body, nil
}
