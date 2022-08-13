package core

import (
	"github.com/ser163/WordBot/generate"
)

// GetRand 获取随机数
func GetRand() string {
	gxa, err := generate.GenRandomWorld(5, "mix")
	if err != nil {
		return "ZheTian"
	}
	return gxa.Word
}
