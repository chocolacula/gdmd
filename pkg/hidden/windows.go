//go:build windows

package hidden

import (
	"log/slog"
	"path/filepath"
	"syscall"
)

func Hidden(path string) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		slog.Warn(err.Error())
		return false
	}

	// https://docs.microsoft.com/en-us/windows/win32/fileio/maximum-file-path-limitation?tabs=cmd
	ptrPath, err := syscall.UTF16PtrFromString(`\\?\` + absPath)
	if err != nil {
		slog.Warn(err.Error())
		return false
	}

	attr, err := syscall.GetFileAttributes(ptrPath)
	if err != nil {
		slog.Warn(err.Error())
		return false
	}

	return (attr & syscall.FILE_ATTRIBUTE_HIDDEN) != 0
}
