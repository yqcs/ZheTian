package server

import (
	"github.com/ser163/WordBot/generate"
)

func GetRand() (w string, er error) {

	gxa, err := generate.GenRandomWorld(5, "mix")
	if err != nil {
		return gxa.Word, err
	}

	return gxa.Word, err
}
