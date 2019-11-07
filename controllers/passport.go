package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yangxinwei/gin-demo-api/types"
)

type Passport struct {
	*BaseController
}

// @Summary 注册
// @Description 注册
// @Accept  json
// @Produce  json
// @Param   mobile     query    string     true         "手机号码"
// @Param   code     query    string     true        	"短信验证码"
// @Param   deviceId     query    string     true        	"设备id"
// @Success 200 {object} types.RegisterResp	"成功"
// @Failure 100101001 {object} types.RegisterResp	"手机号不能为空"
// @Failure 100102001 {object} types.RegisterResp	"手机验证码不能为空"
// @Router /passport/register [POST]
func (c *Passport) Register(ctx *gin.Context) {
	resp := new(types.RegisterResp)

	c.Output(ctx, "200", resp.Data)
}

// @Summary 用户登录
// @Description 用户登录
// @Accept  json
// @Produce  json
// @Param   mobile     query    string     true         "手机号码"
// @Param   code     query    string     true        	"短信验证码"
// @Success 200 {object} types.LoginResp	"成功"
// @Failure 100100001 {object} types.LoginResp	"手机号不能为空"
// @Router /passport/login [POST]
func (c *Passport) Login(ctx *gin.Context) {
	resp := new(types.LoginResp)

	c.Output(ctx, "200", resp.Data)
}

// @Summary 短信获取验证码
// @Description 短信获取验证码
// @Accept  json
// @Produce  json
// @Param   mobile     query    string     true         "手机号码"
// @Success 200 {object} types.SMSCodeResp	"成功"
// @Failure 100100001 {object} types.LoginResp	"手机号不能为空"
// @Router /passport/smscode [POST]
func (c *Passport) SMSCode(ctx *gin.Context) {
	resp := new(types.SMSCodeResp)

	c.Output(ctx, "200", resp.Data)
}

// @Summary 获取用户信息
// @Description 通过deviceId或sessionId获取用户信息
// @Accept  json
// @Produce  json
// @Param   mobile     query    string     true         "手机号码"
// @Success 200 {object} types.GetUserInfoResp	"成功"
// @Failure 100100001 {object} types.LoginResp	"手机号不能为空"
// @Router /passport/getUserInfo [POST]
func (c *Passport) GetUserInfo(ctx *gin.Context) {
	resp := new(types.GetUserInfoResp)

	c.Output(ctx, "200", resp.Data)
}

// @Summary 修改手机号
// @Description 修改手机号
// @Accept  json
// @Produce  json
// @Param   mobile     query    string     true         "mobile number"
// @Param   code     query    string     true        	"短信验证码"
// @Success 200 {object} types.ChangeMobileResp	"成功"
// @Failure 100100001 {object} types.LoginResp	"手机号不能为空"
// @Router /passport/changeMobile [POST]
func (c *Passport) ChangeMobile(ctx *gin.Context) {
	resp := new(types.ChangeMobileResp)

	c.Output(ctx, "200", resp.Data)
}

// @Summary 修改密码
// @Description 修改密码
// @Accept  json
// @Produce  json
// @Param   sessionId     query    string     true         "会话id"
// @Param   new     query    string     true         "新密码"
// @Param   code     query    string     true        	"短信验证码"
// @Success 200 {object} types.ChangePasswordResp	"成功"
// @Failure 100100001 {object} types.ChangePasswordResp	"短信验证码不能为空"
// @Router /passport/changePassword [POST]
func (c *Passport) ChangePassword(ctx *gin.Context) {
	resp := new(types.ChangePasswordResp)

	c.Output(ctx, "200", resp.Data)
}

// @Summary 退出
// @Description 退出、清除session
// @Accept  json
// @Produce  json
// @Success 200 {object} types.LogoutResp	"成功"
// @Router /passport/logout [POST]
func (c *Passport) Logout(ctx *gin.Context) {
	resp := new(types.LogoutResp)

	c.Output(ctx, "200", resp.Data)
}
