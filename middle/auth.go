package middle

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/r2day/base/conf"
	"github.com/r2day/base/log"
)

// CheckCookieIfJWTOK demo how to run a middle for route groups
func CheckCookieIfJWTOK() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Logger.Println("before check the CheckCookieIfJWTOK -->", c.Params)
		//cookie, err := c.Cookie("jwt")
		//if err != nil {
		//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid param"})
		//	return
		//}
		//logger.Logger.Println("second 2 check the CheckCookieIfJWTOK")
		//// 验证cookie
		//token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		//	return []byte(common.ConfInstance.SecretKey), nil
		//})
		//if err != nil {
		//	logger.Logger.WithField("call", "ParseWithClaims").Error(err)
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no auth"})
		//	return
		//}
		//claims := token.Claims.(*jwt.StandardClaims)
		//logger.Logger.Println(claims)
		//var user model.User
		//model.DataHandler.Where("id = ?", claims.Issuer).First(&user)
		//if user.Id == 0 {
		//	logger.Logger.WithField("query", "userInfo")
		//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "no found"})
		//	return
		//}

		//// Set example variable
		//c.Set("user_id",  user.Id)
		//c.Set("user_name", user.Name)
		c.Set("user_id", 1001)
		c.Set("user_name", "frank")
		//logger.Logger.Println("set user info to header successful")
		c.Next()

		// after request
		latency := time.Since(t)
		log.Logger.Print(latency)

		// access the status we are sending
		//status := c.Writer.Status()
		//logger.Logger.Println("the status-->", status)
		//logger.Logger.Println("after check the CheckCookieIfJWTOK")
	}
}

// CORSMiddleware 跨站请求
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//logger.Logger.Println("request params  -->", c.Params)
		c.Writer.Header().Set("Access-Control-Allow-Origin", conf.ConfInstance.AllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
