package im

import (
	"encoding/json"
	"errors"
)

const adminBaseUrl = "https://api.im.jpush.cn/v1/admins/"

type Admin struct {
	IM
}

func NewAdmin(config *Config) *Admin {
	return &Admin{
		IM{config: config},
	}
}

// 用户注册   这里为啥不传UserInfo 是因为那边有些参数存在无值会报错
func (i *Admin) Register(userList map[string]interface{}) (error, UserInfoResult) {
	result := UserInfoResult{}
	i.setUri(adminBaseUrl)
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
