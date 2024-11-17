//go:build !windows

package hidden

import "path/filepath"

func Hidden(path string) bool {
	path = filepath.Base(path)

	return len(path) != 0 && path[0] == '.'
}
