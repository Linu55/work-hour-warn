package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"work-hour-warn/conf"
	"work-hour-warn/dao"
	"work-hour-warn/enity"
	"work-hour-warn/utils"
)

func setupRouter() *gin.Engine {
	conf.InitMysql()
	r := gin.Default()
	r.GET("/lazyBoys", func(c *gin.Context) {
		today, isWorkday := utils.GetWorkday(time.Now(), 1)

		if !isWorkday {
			c.JSON(http.StatusOK, gin.H{"msg": "今天为非工作日，无需查询以及通报"})
		} else {
			dayType, err := strconv.Atoi(c.Query("type"))
			if err != nil {
				fmt.Printf("err=%v", err)
			}
			if dayType != 0 && dayType != 1 {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "type参数输入有误"})
			} else {
				usernames := c.Query("users")
				var groupusernames []string
				groupusernames = strings.Split(usernames, ",")

				uns := conf.GetGroup()
				subtr := utils.Subtr(groupusernames, uns) //获取到请求参数中错误的小组成员
				if len(subtr) != 0 {
					c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数中包含错误的小组成员"})
				} else {
					var unss string
					for _, v := range uns {
						unss = unss + v + "|"
					}
					unss = strings.TrimSuffix(unss, "|") //请求参数tousers

					var content string //请求参数
					var APIDatas []enity.APIData

					var gms []enity.GroupMember

					if dayType == 1 {
						content = ">今日工时填写情况：  \n"
						APIDatas, gms = dao.GetLazyGuy(uns, today)
					} else {
						content = ">上个工作日工时填写情况：  \n"
						lastWorkday, _ := utils.GetWorkday(time.Now(), 0)
						APIDatas, gms = dao.GetLazyGuy(uns, lastWorkday)
					}

					members := make(map[string]string)
					for _, g := range gms {
						members[g.Username] = g.NickName
					}

					for _, v1 := range uns {
						i := 0
						for _, v2 := range APIDatas {
							if v1 == v2.Username {
								if v2.WorkHour < 7.50 {
									var ctt string
									if dayType == 1 {
										ctt = fmt.Sprintf("><font color=\"warning\">DANGER</font> %s: %g  \n", v2.NickName, v2.WorkHour)
									} else {
										ctt = fmt.Sprintf("><font color=\"warning\">DANGER</font> %s: %g  \n", v2.NickName, v2.WorkHour)
									}
									content = content + ctt
								} else {
									var ctt1 string
									if dayType == 1 {
										ctt1 = fmt.Sprintf("><font color=\"info\">OK</font> %s: %g  \n", v2.NickName, v2.WorkHour)
									} else {
										ctt1 = fmt.Sprintf("><font color=\"info\">OK</font> %s: %g  \n", v2.NickName, v2.WorkHour)
									}
									content = content + ctt1
								}
								i = 1
							}
						}
						if i == 1 {
							continue
						}
						var ctt2 string
						if dayType == 1 {
							ctt2 = fmt.Sprintf("><font color=\"warning\">DANGER</font> %s: 0  \n", members[v1])
						} else {
							ctt2 = fmt.Sprintf("><font color=\"warning\">DANGER</font> %s: 0  \n", members[v1])
						}
						content = content + ctt2
					}
					content1 := map[string]string{"content": content}

					println(content)

					m := make(map[string]interface{})
					m["msgtype"] = "markdown"
					m["safe"] = 0
					m["markdown"] = content1
					m["touser"] = unss
					marshal, err := json.Marshal(m)
					if err != nil {
						return
					}
					reader := bytes.NewReader(marshal)
					request, err1 := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=b0ea761c-e403-4465-82e8-038b8e6cd322", reader)
					if err1 != nil {
						return
					}
					request.Header.Set("Content-Type", "application/json")
					client := &http.Client{}
					response, err2 := client.Do(request)
					if err2 != nil {
						return
					}
					defer response.Body.Close()
					body, err3 := ioutil.ReadAll(response.Body)
					if err3 != nil {
						return
					}
					fmt.Printf(string(body))
				}
			}

		}
	})

	return r

}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
