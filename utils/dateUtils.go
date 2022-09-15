package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetWorkday(t time.Time, i int) (string, bool) {
	var day time.Time
	if i == 0 { //如果i==0，获取上一个工作日
		day = t.AddDate(0, 0, -1)
	} else if i == 1 { //如果i==1，判断今天是否为工作日
		day = t
	}
	Year := strconv.Itoa(day.Year())
	Month := strconv.Itoa(int(day.Month()))
	Day := strconv.Itoa(day.Day())
	if int(day.Month()) < 10 {
		Month = "0" + Month
	}
	if day.Day() < 10 {
		Day = "0" + Day
	}
	dateTime := Year + "-" + Month + "-" + Day
	response, err := http.Get("http://timor.tech/api/holiday/info/" + dateTime)
	if err != nil {
		println("err=", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	str := string(body)
	var tempMap map[string]interface{}
	var dayType map[string]interface{}
	err1 := json.Unmarshal([]byte(str), &tempMap)
	if err1 != nil {
		println("err1=", err1)
	}
	for key, value := range tempMap { //获取到集合中type参数
		if key == "type" {
			dayType = value.(map[string]interface{})
		}
	}
	if i == 0 { //上一个工作日为非工作日，则再获取上上一天
		//对type参数进行判断
		if tempMap["holiday"] != nil || dayType["week"].(float64) > 5.0 {
			return GetWorkday(day, i)
		}
	} else { //判断今天为非工作日则返回false
		if tempMap["holiday"] != nil || dayType["week"].(float64) > 5.0 {
			return dateTime, false
		}
	}
	return dateTime, true
}
