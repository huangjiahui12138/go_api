package middlewares

import (
	"jzapi/library"
	"fmt"
	// "time"
	"github.com/gin-gonic/gin"
	"net/url"
)

//校验sign得中间件
func SignMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件开始了")
		var method = c.Request.Method
		// var ts int64
		var sign string
		var req url.Values
		if method == "POST" {
			c.Request.ParseForm()
			req = c.Request.PostForm
		    sign = c.PostForm("sign")
			// ts, _ = strconv.ParseInt(c.PostForm("ts"), 10, 64)
		} else if method == "GET" {
			req = c.Request.URL.Query()
			// sign = c.Query("sign")
		}
		fmt.Printf("%v\n", req)
		if sign == "" || sign != library.CreateSign(req, "123456") {
			fmt.Println("校验失败")
			c.Abort()
			return
		}
		c.Next()
		fmt.Println("中间件结束")
	}
}