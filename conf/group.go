package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetGroup() []string {
	file, err := ioutil.ReadFile("D:/go_project/src/work-hour-warn/config.yaml")
	if err != nil {
		println("err=", err)
	}
	a := make(map[string]interface{})
	err1 := yaml.Unmarshal(file, a)
	if err1 != nil {
		println("err=", err1)
	}
	b := a["group"].(map[interface{}]interface{})
	var usernames []string
	for _, v := range b {
		usernames = append(usernames, v.(string))
	}
	return usernames
}
