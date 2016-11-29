package serial

import (
	"syscall"
	"testing"
	"unsafe"
)

//getTimeout retrive the old timeout settings to see what the last serial controller used
func getTimeout(address string) (c_COMMTIMEOUTS, error) {
	c := &Config{Address: address}
	var timeouts c_COMMTIMEOUTS
	handle, err := newHandle(c)
	if err != nil {
		return timeouts, err
	}
	getCommTimeouts := modkernel32.NewProc("GetCommTimeouts")
	r1, _, e1 := syscall.Syscall(getCommTimeouts.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(&timeouts)), 0)
	if r1 == 0 {
		if e1 != 0 {
			return timeouts, error(e1)
		} else {
			return timeouts, syscall.EINVAL
		}
	}
	return timeouts, nil
}

func TestGetTimeout(t *testing.T) {
	address := "com3"
	timeouts, err := getTimeout(address)
	if err != nil {
		t.Skipf("%v %v", address, err)
	}
	t.Logf("%+v\n", timeouts)

	//sourceforge.net/projects/qmodmaster/ 0.4.6
	//{ReadIntervalTimeout:500 ReadTotalTimeoutMultiplier:0 ReadTotalTimeoutConstant:500 WriteTotalTimeoutMultiplier:0 WriteTotalTimeoutConstant:1000}
	//
	//calta.com Mdbus 3.40
	//slave: {ReadIntervalTimeout:0 ReadTotalTimeoutMultiplier:0 ReadTotalTimeoutConstant:0 WriteTotalTimeoutMultiplier:0 WriteTotalTimeoutConstant:0}
	//master with 2 sec timeout: {ReadIntervalTimeout:0 ReadTotalTimeoutMultiplier:0 ReadTotalTimeoutConstant:2000 WriteTotalTimeoutMultiplier:0 WriteTotalTimeoutConstant:0}

}
