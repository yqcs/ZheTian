![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/logo.png "my-logo")



[# ZheTian](https://github.com/yqcs/ZheTian/)
 

####ZheTian Powerful remote load and execute ShellCode tool
免杀cs shellcode加载框架
>ShellCode加载类型：
-r 读取本地文件内shellcode
-u 从http服务器获取动态shellcode

### 注意：必须使用base64加密，否则无法解析！！
后端使用了base64对读取到的内容进行转码，为了减小流量特征。AES/DES的加密模式太多在未标明的情况下可能导致无法解密失败，故此暂不考虑使用
![my-logo.png](https://raw.githubusercontent.com/yqcs/ZheTian/master/images/1.png "my-logo")
