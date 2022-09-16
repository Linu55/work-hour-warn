package enity

import (
	"fmt"
	"strings"
	"work-hour-warn/conf"
)

type ResponseData struct {
	ToUsers string
	Content string
}

func (r *ResponseData) GetResData(resData *RequestData) *ResponseData {
	uns := conf.GetGroup()
	var unss string
	for _, v := range uns {
		unss = unss + v + "|"
	}
	unss = strings.TrimSuffix(unss, "|") //请求参数tousers

	var ctt string //请求参数
	if resData.DayType == 1 {
		ctt = ">今日工时填写情况：  \n"
	} else {
		ctt = ">上个工作日工时填写情况：  \n"
	}

	members := make(map[string]string)
	for _, g := range resData.Users {
		members[g.Username] = g.NickName
	}

	for _, v1 := range uns {
		i := 0
		for _, v2 := range resData.UserDetail {
			if v1 == v2.Username {
				if v2.WorkHour < 7.50 {
					var ctt1 string
					ctt1 = fmt.Sprintf("><font color=\"warning\">DANGER</font> %s: %g  \n", v2.NickName, v2.WorkHour)
					ctt = ctt + ctt1
				} else {
					var ctt2 string
					ctt2 = fmt.Sprintf("><font color=\"info\">OK</font> %s: %g  \n", v2.NickName, v2.WorkHour)
					ctt = ctt + ctt2
				}
				i = 1
			}
		}
		if i == 1 {
			continue
		}
		var ctt3 string
		ctt3 = fmt.Sprintf("><font color=\"warning\">DANGER</font> %s: 0  \n", members[v1])
		ctt = ctt + ctt3
	}

	var res ResponseData
	res.ToUsers = unss
	res.Content = ctt
	return &res

}
