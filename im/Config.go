package im

import "encoding/base64"

const (
	BASE64TABLE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

type Config struct {
	appKey       string
	masterSecret string
}

func (c *Config) SetConfig(AppKey, MasterSecret string) {
	c.appKey = AppKey
	c.masterSecret = MasterSecret
}

var base64Coder = base64.NewEncoding(BASE64TABLE)

func (c *Config) getAuthorization() string {
	return "Basic " + base64Coder.EncodeToString([]byte(c.appKey+":"+c.masterSecret))
}

func NewConfig() *Config {
	return &Config{}
}
