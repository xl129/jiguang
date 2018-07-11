package im

import (
	"errors"
	"encoding/json"
	"fmt"
)

const messagesBaseUrl = "https://api.im.jpush.cn/v1/messages"

type Messages struct {
	IM
	MessagesContent
}

type MessagesContent struct {
	Version        int8   `json:"version"`         //版本号 目前是1 （必填）
	TargetType     string `json:"target_type"`     //发送目标类型 single - 个人，group - 群组 chatroom - 聊天室（必填）
	FromType       string `json:"from_type"`       //发送消息者身份 当前只限admin用户，必须先注册admin用户 （必填）
	MsgType        string `json:"msg_type"`        //发消息类型 text - 文本，image - 图片, custom - 自定义消息（msg_body为json对象即可，服务端不做校验）voice - 语音 （必填）
	TargetId       string `json:"target_id"`       //目标id single填username group 填Group id chatroom 填chatroomid（必填）
	TargetAppkey   string `json:"target_appkey"`   //跨应用目标appkey（选填）
	FromId         string `json:"from_id"`         //发送者的username （必填
	FromName       string `json:"from_name"`       //发送者展示名（选填）
	TargetName     string `json:"target_name"`     //接受者展示名（选填）
	NoOffline      bool   `json:"no_offline"`      //消息是否离线存储 true或者false，默认为false，表示需要离线存储（选填）
	NoNotification bool   `json:"no_notification"` //消息是否在通知栏展示 true或者false，默认为false，表示在通知栏展示（选填）
	Notification struct {
		Title string `json:"title"` //通知的标题（选填）
		Alert string `json:"alert"` //通知的内容（选填）
	} `json:"notification"`                        //自定义通知栏展示（选填）
	MsgBody interface{} `json:"msg_body"`          //Json对象的消息体 限制为4096byte
}

type MessagesResult struct {
	MsgId    int64 `json:"msg_id"`
	MsgCtime int64 `json:"msg_ctime"`
	Error          `json:"error"`
}

func NewMessages(config *Config) *Messages {
	return &Messages{
		IM{config: config},
		MessagesContent{Version: 1},
	}
}

// 设置版本号
func (i *Messages) SetVersion() {
	i.MessagesContent.Version = 1
}

// 设置发送目标类型
func (i *Messages) SetTargetType(targetType string) error {
	t := [3]string{"single", "group", "chatroom"}
	check := false
	for _, v := range t {
		if v == targetType {
			check = true
			break
		}
	}
	if !check {
		return errors.New("发送目标类型targetType错误")
	}
	i.MessagesContent.TargetType = targetType
	return nil
}

// 发送消息者身份
func (i *Messages) SetFromType(fromType string) {
	i.MessagesContent.FromType = fromType
}

// 发消息类型
func (i *Messages) SetMsgType(msgType string) error {
	t := [4]string{"text", "image", "custom", "voice"}
	check := false
	for _, v := range t {
		if v == msgType {
			check = true
			break
		}
	}
	if !check {
		return errors.New("发消息类型msgType错误")
	}
	i.MessagesContent.MsgType = msgType
	return nil
}

// 目标id single填username group 填Group id chatroom 填chatroomid（必填）
func (i *Messages) SetTargetId(targetId string) {
	i.MessagesContent.TargetId = targetId
}

// 跨应用目标appkey
func (i *Messages) SetTargetAppkey(targetAppkey string) {
	i.MessagesContent.TargetAppkey = targetAppkey
}

//发送者的username
func (i *Messages) SetFromId(fromId string) {
	i.MessagesContent.FromId = fromId
}

//发送者展示名
func (i *Messages) SetFromName(fromName string) {
	i.MessagesContent.FromName = fromName
}

//接受者展示名
func (i *Messages) SetTargetName(targetName string) {
	i.MessagesContent.TargetName = targetName
}

//消息是否离线存储
func (i *Messages) SetNoOffline(b bool) {
	i.MessagesContent.NoOffline = b
}

//消息是否在通知栏展示
func (i *Messages) SetNoNotification(b bool) {
	i.MessagesContent.NoNotification = b
}

//通知的标题（选填）
func (i *Messages) SetNotificationTitle(title string) {
	i.MessagesContent.Notification.Title = title
}

// 通知的内容（选填）
func (i *Messages) SetNotificationAlert(alert string) {
	i.MessagesContent.Notification.Alert = alert
}

// Json对象的消息体 限制为4096byte
func (i *Messages) SetMsgBody(s interface{}) {
	i.MessagesContent.MsgBody = s
}

// 设置文本消息
func (i *Messages) SetTextBody(t string, extras interface{}) {
	body := struct {
		Text   string      `json:"text"`
		Extras interface{} `json:"extras"`
	}{Text: t, Extras: extras}

	i.SetMsgBody(body)
}

// 设置图片消息
func (i *Messages) SetImageBody(image ImageResource) {
	i.SetMsgBody(image)
}

// 设置语音消息
func (i *Messages) SetVoiceBody(voice VoiceResource) {
	i.SetMsgBody(voice)
}

// 自定义的消息格式
func (i *Messages) SetCustom(custom interface{}) {
	i.SetMsgBody(custom)
}

// 参数简单检查
func (i *Messages) check() error {
	if i.MessagesContent.Version != 1 {
		return errors.New("Version必填")
	}
	if i.MessagesContent.TargetType == "" {
		return errors.New("TargetType必填")
	}
	if i.MessagesContent.FromType == "" {
		return errors.New("FromType必填")
	}
	if i.MessagesContent.MsgType == "" {
		return errors.New("MsgType必填")
	}
	if i.MessagesContent.TargetId == "" {
		return errors.New("TargetId必填")
	}
	if i.MessagesContent.FromName == "" {
		return errors.New("FromName必填")
	}
	return nil
}

// 发送消息
func (i *Messages) Send() (error, MessagesResult) {
	result := MessagesResult{}
	err := i.check()
	if err != nil {
		return err, result
	}
	i.setUri(messagesBaseUrl)
	bytes, err := json.Marshal(i.MessagesContent)
	if err != nil {
		return err, result
	}
	fmt.Println(string(bytes))
	Response := i.post(bytes)
	if Response.Err != nil {
		return Response.Err, result
	}
	err = json.Unmarshal(Response.Data, &result)
	if err == nil {
		return err, result
	}
	fmt.Println(Response)
	fmt.Println(string(Response.Data))
	//创建失败，具体信息看返回
	if Response.StatusCode != 201 {
		return errors.New("发送失败"), result
	}
	return nil, result
}
