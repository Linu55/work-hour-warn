package enity

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"work-hour-warn/conf"
	"work-hour-warn/dao"
	"work-hour-warn/pojo"
	"work-hour-warn/utils"
)

var aPIDatas []pojo.APIData
var gms []pojo.GroupMember

type RequestData struct {
	DayType    int
	UserDetail []pojo.APIData
	Users      []pojo.GroupMember
}

func (r *RequestData) GetRequestData(c *gin.Context) (*RequestData, int) {
	dayType, err := strconv.Atoi(c.Query("type"))
	if err != nil {
		fmt.Printf("err=%v", err)
	}
	if dayType != 0 && dayType != 1 {
		return nil, 0
	} else {
		usernames := c.Query("users")
		var groupusernames []string
		groupusernames = strings.Split(usernames, ",")

		uns := conf.GetGroup()
		subtr := utils.Subtr(groupusernames, uns) //获取到请求参数中错误的小组成员
		if len(subtr) != 0 {
			return nil, 1
		} else {
			if dayType == 1 {
				today, _ := utils.GetWorkday(time.Now(), 1)
				aPIDatas, gms = dao.GetLazyGuy(uns, today)
			} else {
				lastWorkday, _ := utils.GetWorkday(time.Now(), 0)
				aPIDatas, gms = dao.GetLazyGuy(uns, lastWorkday)
			}
		}
	}

	var requestData RequestData
	requestData.DayType = dayType
	requestData.UserDetail = aPIDatas
	requestData.Users = gms

	return &requestData, 2
}
