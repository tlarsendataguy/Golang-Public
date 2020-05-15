package convert_strings

import (
	"syscall"
	"unsafe"
)

func StringToWideC(value string) (unsafe.Pointer, error) {
	utf16Bytes, err := syscall.UTF16FromString(value)
	if err != nil {
		return nil, err
	}

	utf16Bytes = append(utf16Bytes, 0)
	return unsafe.Pointer(&utf16Bytes[0]), nil
}

func CToString(char unsafe.Pointer) string {
	if uintptr(char) == 0x0 {
		return ``
	}

	wcharPtr := uintptr(char)
	ws := make([]byte, 0)
	for {
		w := *((*byte)(unsafe.Pointer(wcharPtr)))

		// check if the current wchar is nil and also the first wchar in a UTF-16 sequence.  If yes, we
		// have reached the end of the string
		if w == 0 {
			break
		}
		ws = append(ws, w)

		wcharPtr += 1
	}
	return string(ws)
}

func WideCToString(wchar_t unsafe.Pointer) string {
	if uintptr(wchar_t) == 0x0 {
		return ``
	}

	wcharPtr := uintptr(wchar_t)
	ws := make([]uint16, 0)
	index := 1
	for {
		w := *((*uint16)(unsafe.Pointer(wcharPtr)))

		// check if the current wchar is nil and also the first wchar in a UTF-16 sequence.  If yes, we
		// have reached the end of the string
		if index%2 != 0 && w == 0 {
			break
		}
		ws = append(ws, w)

		wcharPtr += 2
		index += 1
	}
	return syscall.UTF16ToString(ws)
}