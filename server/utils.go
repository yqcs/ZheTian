package server

import (
	"github.com/ser163/WordBot/generate"
)

//GetRand 获取随机数
func GetRand() (w string, er error) {

	gxa, err := generate.GenRandomWorld(5, "mix")
	if err != nil {
		return gxa.Word, err
	}

	return gxa.Word, err
}
