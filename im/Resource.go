package im

import (
	"os"
	"errors"
	"strings"
	"path"
	"encoding/json"
)

//注：文件大小限制8m，暂时只支持图片格式 jpg bmp gif png等
const resourceBaseUrl = "https://api.im.jpush.cn/v1/resource"

var ResourceFileExt map[string][]string

type ResourceResponse struct {
	MediaId    string `json:"media_id"`
	MediaCrc32 int64  `json:"media_crc32"`
	Width      int    `json:"width"` //图片才有
	Height     int    `json:"height"`
	Format     string `json:"format"`
	FSize      int    `json:"fsize"`
	Hash       string `json:"string"`
	FName      string `json:"fname"` //文件才有
	Error             `json:"error"`
}

type Resource struct {
	IM
}

func NewResource(config *Config) *Resource {
	return &Resource{
		IM{config: config},
	}
}

func init() {
	ResourceFileExt = map[string][]string{
		"image": {
			".jpg",
			".gif",
			".png",
			".bmp",
		},
		"voice": {
			".m4a",
		},
		"file": {
			".dmp",
		},
	}
}

//是否是允许上传的文件类型
func (i *Resource) isAllowFileType(fileType string) bool {
	switch fileType {
	case "image":
		return true
	case "voice":
		return true
	case "file":
		return true
	}
	return false
}

//是否是允许的扩展名
func (i *Resource) isAllowExt(fileType string, ext string) bool {
	extArr, ok := ResourceFileExt[fileType]
	if !ok {
		return false
	}
	for _, v := range extArr {
		if ext == strings.ToLower(v) {
			return true
		}
	}
	return false
}
func (i *Resource) checkFile(fileType, filePath string) error {
	if !i.isAllowFileType(fileType) {
		return errors.New("文件类型错误")
	}

	if !i.isAllowExt(fileType, path.Ext(filePath)) {
		return errors.New("文件错误")
	}

	return nil
}

// 上传文件
func (i *Resource) Upload(fileType, path string) (error, ResourceResponse) {
	result := ResourceResponse{}
	err := i.checkFile(fileType, path)
	if err != nil {
		return err, result
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err, result
	}
	if fileInfo.Size() > 8388608 {
		return errors.New("文件超过8M了"), result
	}
	i.setUri(resourceBaseUrl + "?type=" + fileType)
	Response := i.upload(path)
	if Response.Err != nil {
		return Response.Err, result
	}
	err = json.Unmarshal(Response.Data, &result)
	if err == nil {
		return err, result
	}
	if Response.StatusCode != 200 {
		return errors.New("上传失败"), result
	}
	return nil, result
}
