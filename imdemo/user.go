package imdemo

import (
	"jg/im"
	"fmt"
)

func UserRegister() {
	var UserList = make([]map[string]interface{}, 1)
	item := map[string]interface{}{
		"username": "Test6",
		"appkey":   im.AppKey, //用户所在的appkey  估计android ios不能使同一个应用
		"password": "Test",    //极光会md5一次
		"gender":   0,
	}

	UserList[0] = item
	var u = im.NewUser()
	err, data := u.Register(UserList)
	fmt.Println(err)
	fmt.Println(data)
}

func UserUpload() {
	var username = "Test4"
	item := map[string]interface{}{
		"password": "Test1234", //极光会md5一次
		"gender":   1,
	}
	var u = im.NewUser()
	err := u.Update(username, item)
	fmt.Println(err)
}
