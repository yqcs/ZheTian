![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/logo.png "my-logo")



[# ZheTian](https://github.com/yqcs/ZheTian/)
 

####ZheTian Powerful remote load and execute ShellCode tool
免杀cs shellcode加载框架
>ShellCode加载类型：
-r 读取本地文件内shellcode字节码
-u 从http服务器获取动态shellcode字节码
-s 直接加载原始Payload文件

从http远程服务器或本地文件内读取shellcode字节码需使用以下格式： java类型需去除0x，c or python 类型需去除\x

###//示例：
 
  原始python类型： 
>buf = "\xfc\x48\x83\xe4\xf0\xe8\xc8\x00\x00\x00\x41\x51\x41\x50\x52\x51\x56\x48\x31\xd2\x65\x48\x8b"....
 
修改成：
>fc4883e4f0e8c8000000415141......
> 
Java、c、ruby等同理。只需将字节数组里的16进制代码提取出来即可。暂时只支持64位payload

### 注意：使用code字节码加载时必须使用base64加密，否则无法解析！！
后端使用了base64对读取到的内容进行转码，为了减小流量特征。AES/DES的加密模式太多在未标明的情况下可能导致无法解密失败，故此暂不考虑使用
![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/1.png "my-logo")
