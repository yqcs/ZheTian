package dev

import (
	"fmt"
	"syscall"
	"unsafe"
)

type (
	LPVOID         uintptr
	LMSTR          *uint16
	DWORD          uint32
	LPBYTE         *byte
	LPDWORD        *uint32
	LPWSTR         *uint16
	NET_API_STATUS DWORD

	USER_INFO_1 struct {
		Usri1_priv         int // USER_PRIV_GUEST = 0 USER_PRIV_USER  = 1 USER_PRIV_ADMIN = 2
		Usri1_name         string
		Usri1_password     string
		Usri1_password_age int
		Usri1_home_dir     string
		Usri1_comment      string
		Usri1_flags        int
		Usri1_script_path  string
	}
)

var (
	nodnetapi32    = syscall.NewLazyDLL("netapi32.dll")
	procNetUserAdd = nodnetapi32.NewProc("NetUserAdd")
)

func NetUserAdd(servername LMSTR, level DWORD, buf LPBYTE, parm_err LPDWORD) (status NET_API_STATUS) {

	r0, _, _ := syscall.Syscall6(procNetUserAdd.Addr(), 4, uintptr(0), uintptr(level), uintptr(unsafe.Pointer(&buf)), uintptr(0), 0, 0)
	status = NET_API_STATUS(r0)
	fmt.Println(status)
	return
}

//Adduser 添加系统用户
//TODO::
func Adduser() {
	var newUser USER_INFO_1
	newUser.Usri1_priv = 1
	//newUser.Usri1_password_age=1
	//newUser.Usri1_flags = 1
	newUser.Usri1_name = "xs@sssssad12"
	newUser.Usri1_password = "xxadasd1@@ass"
	newUser.Usri1_comment = "测试系统"
	data := (*byte)(unsafe.Pointer(&newUser))
	NetUserAdd(nil, 1, data, nil)
}
