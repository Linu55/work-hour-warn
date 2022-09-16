package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"work-hour-warn/enity"
	"work-hour-warn/utils"
	"work-hour-warn/wxinterface"
)

// Warn godoc
// @Summary 工时预警接口
// @version 1.0
// @description  通过请求参数，通知指定人员的今日或者上个工作日工时填写情况。
// @Host 127.0.0.1:8080
// @accept json
// @Produce  json
// @Param type query string true "查询类型 0：查询上个工作日工时填写情况；1：查询今日工时填写情况"
// @Param users query string true "待查询小组成员 由成员拼音组成，成员之间用|隔开"
// @Success 200 {string} json "{"code":200,"msg":"微信接口请求成功"}"
// @Failure 400 {string} json {"code":400,"msg":"type参数输入有误"}
// @Failure 600 {string} json {"code":600,"msg":"请求参数中包含错误的小组成员或成员名有误"}
// @Success 700 {string} json "{"code":700,"msg":"今天为非工作日，无需查询以及通报"}"
// @Success 800 {string} json "{"code":800,"msg":"微信接口请求失败"}"
// @Router /lazyBoys [get]
func Warn(c *gin.Context) {
	_, isWorkday := utils.GetWorkday(time.Now(), 1)
	if !isWorkday {
		c.JSON(700, gin.H{"msg": "今天为非工作日，无需查询以及通报"})
	} else {
		requestData := new(enity.RequestData)
		reqd, code := requestData.GetRequestData(c)
		if code == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "type参数输入有误"})
		} else {
			if code == 1 {
				c.JSON(600, gin.H{"msg": "请求参数中包含错误的小组成员"})
			} else {
				responseData := new(enity.ResponseData)
				resd := responseData.GetResData(reqd)

				rbt := new(wxinterface.RbtData)
				ok := rbt.SendRequest(resd)
				if ok {
					c.JSON(http.StatusOK, gin.H{"msg": "微信接口请求成功"})
				} else {
					c.JSON(800, gin.H{"msg": "微信接口请求失败"})
				}
			}
		}

	}
}
