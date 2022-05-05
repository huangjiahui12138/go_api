package sdk

import (
	"github.com/gin-gonic/gin"
)

type Device struct {

}


func (d *Device) Init(ctx *gin.Context) {
	// var gid = ctx.PostForm("gid")
	// var gkey = ctx.PostForm("gkey")
	// var deviceno = ctx.PostForm("deviceno")
	// var device_type = ctx.PostForm("type")
	// var operator = ctx.PostForm("operator")       // sim卡所属运营商
	// var phoneModel = ctx.PostForm("phoneModel")   //设备型号
	// var screenSize = ctx.PostForm("screenSize")    // 屏幕大小
	// var ctype = ctx.DefaultPostForm("ctype", "own")
	// var imei = ctx.PostForm("imei")
	// var sysOs = ctx.PostForm("sysOs")              // 设备系统版本型号
	// var androidid = ctx.PostForm("androidid")
	// var oaid = ctx.PostForm("oaid")
	// var pk_ts = ctx.PostForm("pk_ts")

	// if gid == "" {
	// 	ctx.JSON(200, map[string]interface{}{"msg": "gid is null", "status": 300})
	// 	return
	// }

	// if deviceno == "" {

	// }




	ctx.JSON(200, map[string]interface{}{"msg": "success", "status": 200})
	return
}


// func DeviceInit(c *gin.Context) {
// 	name := c.PostForm("name")
// 	//获取所有post参数
// 	// ReqPost := c.Request.PostForm
	
// 	// c.JSON(200, gin.H{
// 	// 	"v1" : "deviceInit",
// 	// 	"name" : "name1",
// 	// })
// 	fmt.Println(123)
// 	c.JSON(200, map[string]interface{}{
// 		"msg" : "ok",
// 		"name" : name,
// 	})
// }