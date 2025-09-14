package files

import (
	"io/fs"
	"os"
)

const (
	// Rw represents a file permission of read/write for current user
	// and no access for user's group and other groups.
	Rw fs.FileMode = 0o600

	// RwRR represents a file permission of read/write for current user
	// and read-only access for user's group and other groups.
	RwRR fs.FileMode = 0o644

	// RwRwRw represents a file permission of read/write for current user
	// and read/write too for user's group and other groups.
	RwRwRw fs.FileMode = 0o666

	// RwxRxRxRx represents a file permission of read/write/execute for current user
	// and read/execute for user's group and other groups.
	RwxRxRxRx fs.FileMode = 0o755
)

// Exists returns a boolean indicating whether the provided input src can be stat'ed or not.
func Exists(src string) bool {
	_, err := os.Stat(src)
	return err == nil
}
