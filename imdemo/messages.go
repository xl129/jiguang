package imdemo

import (
	"jg/im"
	"fmt"
)

func MessageSend() {
	u := im.NewMessages(getConfig())
	u.SetVersion()
	u.SetTargetType("single")
	u.SetFromType("admin")
	u.SetMsgType("text")
	u.SetTargetId("V5O0mTHc4aKnzgfW1nGyLpC4m7JAEY2d")
	u.SetTargetAppkey(appKey)
	u.SetFromId("admin")
	u.SetFromName("系统消息")
	u.SetTargetName("")
	u.SetNoOffline(false)
	u.SetNoNotification(false)
	u.SetNotificationTitle("Test")
	u.SetNotificationAlert("这是一个测试")
	body := struct {
		Text   string   `json:"text"`
		Extras struct{} `json:"extras"`
	}{Text: "Hello"}
	u.SetMsgBody(body)
	err, result := u.Send()
	fmt.Println(err)
	fmt.Println(result)
}
