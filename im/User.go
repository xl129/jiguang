package im

import (
	"encoding/json"
	"errors"
)

const userBaseUrl = "https://api.im.jpush.cn/v1/users/"

type User struct {
	IM
}

type UserInfo struct {
	UserName  string `json:"username"`
	AppKey    string `json:"appkey"`   //用户所在的appkey  估计android ios不能使同一个应用
	Password  string `json:"password"` //极光会md5一次
	NickName  string `json:"nickname"`
	Avatar    string `json:"avatar"`    //需要填上从文件上传接口获得的media_id
	Birthday  string `json:"birthday"`  //"1990-01-24 00:00:00"
	Gender    int8   `json:"gender"`    //性别 0 - 未知， 1 - 男 ，2 - 女
	Signature string `json:"signature"` //用户签名
	Region    string `json:"region"`    //用户所属地区
	Address   string `json:"address"`   //用户详细地址
	CTime     string `json:"ctime"`     //创建"1990-01-24 00:00:00"
	MTime     string `json:"mtime"`     //修改"1990-01-24 00:00:00"
	Extras    string `json:"extras"`    //自定义字段 json 对象
}

type UserInfoResult struct {
	UserName string `json:"username"`
	Error           `json:"error"`
}

func NewUser() *User {
	return &User{}
}

// 用户注册   这里为啥不传UserInfo 是因为那边有些参数存在无值会报错
func (i *User) Register(userList []map[string]interface{}) (error, []UserInfoResult) {
	l := len(userList)
	result := make([]UserInfoResult, l)
	if l == 0 {
		return errors.New("请传要添加的用户列表"), result
	}
	if l > 500 {
		return errors.New("一次最多500"), result
	}
	i.setUri(userBaseUrl)
	bytes, err := json.Marshal(userList)
	if err != nil {
		return err, result
	}
	Response := i.post(bytes)
	if Response.Err != nil {
		return Response.Err, result
	}
	err = json.Unmarshal(Response.Data, &result)
	if err == nil {
		return err, result
	}
	//创建失败，具体信息看返回
	if Response.StatusCode != 201 {
		return errors.New("添加失败"), result
	}
	return nil, result
}

// 更新信息
func (i *User) Update(userName string, info map[string]interface{}) error {
	i.setUri(userBaseUrl + userName)
	bytes, err := json.Marshal(info)
	if err != nil {
		return err
	}
	Response := i.put(bytes)
	if Response.Err != nil {
		return Response.Err
	}
	if Response.StatusCode != 204 {
		return errors.New("更新失败")
	}
	return nil
}
