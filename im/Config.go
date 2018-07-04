package im

import "encoding/base64"

const (
	AppKey       = "xxx"
	MasterSecret = "xxx"
	BASE64TABLE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64Coder = base64.NewEncoding(BASE64TABLE)

func getAuthorization() string {
	return "Basic " + base64Coder.EncodeToString([]byte(AppKey+":"+MasterSecret))
}
