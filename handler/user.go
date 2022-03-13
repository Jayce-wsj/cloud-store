package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	db "workspace/cloud-pan/db/mysql"
	"workspace/cloud-pan/util/logs"
)

type UserReq struct {
	Username string `form:"username" json:"username"`
	Password string
}


//用户注册
func SignupHandler(c *gin.Context) {
	rsp := &Result{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
	req := &UserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logs.Errorf("err:%v", err)
		return
	}

	//插入数据库
	db.Insert().Table("user").Value(db.GetValues().Adds(
		"username", req.Username,
		"password", req.Password,
		"regtime", time.Now().Format("2006-01-02 15:04:05"))).Exec()
}

//用户登录
func SigninHandler(c *gin.Context) {
	rsp := &Result{
		Status: 0,
		Msg:    "登陆成功",
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	req := &UserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logs.Errorf("err:%v", err)
		return
	}
	logs.Debugf("req:%+v", req)

	row, err := db.Select().From("user").Where(fmt.Sprintf("username='%s' and password='%s'", req.Username, req.Password)).QueryOne()
	if err != nil {
		logs.Errorf("err:%v", err)
		return
	}
	logs.Debugf("row:%+v", row)

	if row.Get("username") == "" {
		logs.Debug(111)
		rsp.Status = 500
		rsp.Msg = "用户名或密码错误"
		return
	} else {
		rsp.Data = row
	}
	logs.Debug(c.Request.Host)
	logs.Debug(c.Request.RequestURI)
	//登陆成功，设置cookie
	c.SetCookie("username", req.Username, 36000, "/", c.Request.Host, false, true)
}

//用户登录
func InfoHandler(c *gin.Context) {
	rsp := &Result{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	req := &UserReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		logs.Errorf("err:%v", err)
		return
	}
	logs.Debugf("req:%+v", req)

	rsp.Data, _ = db.Select().From("user").Where(fmt.Sprintf("username='%s'", req.Username)).QueryOne()
}
