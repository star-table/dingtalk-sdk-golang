package sdk

import (
	"github.com/flyingtime/dingtalk-sdk-golang/http"
	"github.com/flyingtime/dingtalk-sdk-golang/json"
)

func RobotSender(webHook string, msg WorkNoticeMsg) {
	reqJson, err := json.ToJson(msg)
	if err != nil {
		panic(err)
	}
	http.Post(webHook, nil, reqJson)
}
