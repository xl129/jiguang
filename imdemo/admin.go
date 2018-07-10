package imdemo

import (
	"fmt"
	"jg/im"
)

func AdminRegister() {
	item := map[string]interface{}{
		"username": "admin",
		"password": "Test", //极光会md5一次
	}

	var u = im.NewAdmin(getConfig())
	err, data := u.Register(item)
	fmt.Println(err)
	fmt.Println(data)
}
