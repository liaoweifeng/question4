package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"question4/model"
	"strconv"
	"time"
)

var daystart, dayend int64

func main() {
	router := gin.Default()
	router.POST("/register", Registe)
	router.POST("/login", Login)
	router.POST("/signin", SignDay)
	router.POST("/operate", Operate)
	router.Run(":8080")
}

//注册
func Registe(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("user:" + username + " " + password)
	if model.UserSignup(username, password) {
		c.JSON(500, gin.H{"status": http.StatusInternalServerError, "message": "数据库报错或者用户名" + username + "已经被注册"})
	} else {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
	}
}

//登录并设定cookie
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if model.UserSignin(username, password) {
		power := model.Checkpower(username)
		power1 := strconv.Itoa(power)
		c.SetCookie("username", username, 100, "/", "localhost", false, true)
		c.SetCookie("power", power1, 100, "/", "localhost", false, true)
		//第一个参数为 cookie 名；第二个参数为 cookie 值；第三个参数为 cookie 有效时长；第四个参数为 cookie 所在的目录；第五个为所在域，表示我们的 cookie 作用范围；第六个表示是否只能通过 https 访问；第七个表示 cookie 是否支持HttpOnly属性。

		c.JSON(200, gin.H{"status": http.StatusOK, "message": "登录成功,在签到页面输入'qiandao'进行签到,一天只能签到一次"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "登录失败，用户名或密码错误"})
	}
}

//签到
func SignDay(c *gin.Context) {
	message := c.PostForm("message")
	username, err := c.Cookie("username")
	fmt.Println("username" + ":" + username)
	if err != nil {
		c.JSON(500, gin.H{"status": http.StatusForbidden, "message": "cookie读取失败（您还没有登录）"})
		return
	}
	daystart = time.Now().Unix()
	fmt.Println(daystart)
	fmt.Println(dayend)
	if daystart-dayend >= 86400 {
		dayend = time.Now().Unix()
		if model.Signday(username, message) {
			c.JSON(200, gin.H{"内容": "签到成功", "用户名": username})
		} else {
			c.JSON(403, gin.H{"status": http.StatusForbidden, "signday": "签到失败"})
		}
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "signday": "签到失败,你签到的时间间隔不足一天"})
	}
}

//管理员管理签到系统
func Operate(c *gin.Context) {
	username, err := c.Cookie("username")
	fmt.Println("username" + ":" + username)
	if err != nil {
		c.JSON(500, gin.H{"status": http.StatusForbidden, "message": "cookie读取失败（您还没有登录）"})
		return
	}
}
