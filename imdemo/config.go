package imdemo

import "jg/im"

const appKey = "xxxx"
const masterSecret = "xxxx"

func getConfig() *im.Config {
	config := im.NewConfig()
	config.SetConfig(appKey, masterSecret)
	return config
}
