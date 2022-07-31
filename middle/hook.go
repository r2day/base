package middle

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/r2day/base/conf"
	logger "github.com/r2day/base/log"
	"github.com/r2day/base/util"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func BeforeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {

		payload, err := c.GetRawData()
		if err != nil {
			logger.Logger.Error(err)
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
		hash := util.SignWithMd5(payload, conf.ConfInstance.SignKey)
		// TODO check hash

		// 请求ID 用于客户需要定位问题时，提供
		requestId := c.Writer.Header().Get("Request-Id")

		// 日志中不应该明文打印payload 防止被拦截获得
		// TODO payload 加密
		encoded := base64.StdEncoding.EncodeToString(payload)

		entry := logger.Logger.WithFields(log.Fields{
			"client_ip":  util.GetClientIP(c),
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    util.GetUserID(c),
			"referrer":   c.Request.Referer(),
			"hash":       hash,
			"payload":    encoded,
			"request_id": requestId,
			// "api_version": util.ApiVersion,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("before handler request")
		}
		// 设置返回头
		c.Header("Request-ID", requestId)

		// Process Request
		c.Next()
	}
}

func AfterRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		// Process Request
		c.Next()

		// Stop timer
		duration := util.GetDurationInMilliseconds(start)
		fmt.Sprintf("url=%s, status=%d, resp=%s", c.Request.URL, c.Writer.Status(), blw.body.String())

		// TODO resp 加密
		encoded := base64.StdEncoding.EncodeToString(blw.body.Bytes())

		// 请求ID 用于客户需要定位问题时，提供
		requestId := c.Writer.Header().Get("Request-Id")

		entry := logger.Logger.WithFields(log.Fields{
			"duration":   duration,
			"request_id": requestId,
			"payload":    encoded,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("after handler request")
		}
		// 设置返回头
		c.Header("Request-ID", requestId)

	}
}
