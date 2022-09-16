package wxinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"work-hour-warn/enity"
)

type RbtData struct {
	MsgType string
}

func (r *RbtData) SendRequest(resd *enity.ResponseData) bool {
	content := map[string]string{"content": resd.Content}
	m := make(map[string]interface{})
	m["msgtype"] = "markdown"
	m["safe"] = 0
	m["markdown"] = content
	m["touser"] = resd.ToUsers
	marshal, err := json.Marshal(m)
	if err != nil {
		return false
	}
	reader := bytes.NewReader(marshal)
	request, err1 := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=b0ea761c-e403-4465-82e8-038b8e6cd322", reader)
	if err1 != nil {
		return false
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err2 := client.Do(request)
	if err2 != nil {
		return false
	}
	defer response.Body.Close()
	body, err3 := ioutil.ReadAll(response.Body)
	if err3 != nil {
		return false
	}
	fmt.Printf(string(body))
	return true
}
