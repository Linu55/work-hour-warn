package test

import (
	"fmt"
	"testing"
	"work-hour-warn/conf"
	"work-hour-warn/utils"
)

func TestGroup(t *testing.T) {
	/*file, err := ioutil.ReadFile("D:/go_project/src/work-hour-warn/config.yaml")
	if err != nil {
		return
	}
	a := make(map[string]interface{})
	err1 := yaml.Unmarshal(file, a)
	if err1 != nil {
		return
	}
	b := a["group"].(map[interface{}]interface{})
	for k, v := range b {
		fmt.Printf("k=%v,v=%v,", k, v)
	}*/
	usernames := conf.GetGroup()
	for i, v := range usernames {
		fmt.Printf("i=%v,v=%v\n", i, v)
	}

	println("**************************************")
	usernames1 := []string{"fengyuangen", "tangjie", "yuguo", "lilin", "chenziang", "xuzhuo"}
	for i, v := range usernames1 {
		fmt.Printf("i=%v,v=%v\n", i, v)
	}

	println("**************************************")

	substr := utils.Subtr(usernames, usernames1)

	fmt.Printf("%v", substr)

}

func Test2(t *testing.T) {

}
