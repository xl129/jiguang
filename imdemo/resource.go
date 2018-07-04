package imdemo

import (
	"jg/im"
	"fmt"
)

func ResourceUpload() {
	var u = im.NewResource()
	err, data := u.Upload("image", "/home/lin/图片/cret.png")
	fmt.Println(err)
	fmt.Println(data)
}
