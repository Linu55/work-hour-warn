package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
)

var SqlSession *gorm.DB
var err2 error

func InitMysql() {
	//将yaml配置参数拼接成连接数据库的url
	file, err := ioutil.ReadFile("D:/go_project/src/work-hour-warn/config.yaml")
	a := map[string]interface{}{}
	if err != nil {
		println("err=", err)
	}
	err1 := yaml.Unmarshal(file, a)
	if err1 != nil {
		println("err1=", err1)
	}
	conf := a["database"].(map[interface{}]interface{})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf["userName"].(string),
		conf["password"].(string),
		conf["url"].(string),
		conf["post"].(int),
		conf["dbname"].(string),
	)
	fmt.Printf("%s\n", dsn)
	//连接数据库
	SqlSession, err2 = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err2 != nil {
		println("err2=", err2)
	}
}
