package files

import (
	"fmt"
	"io"
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

// Copy copies a provided file from src to dest with a default permission of 0o644. It fails if it's a directory.
func Copy(src, dst string) error {
	sfile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer sfile.Close()

	dfile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer dfile.Close()

	// copy buffer from src to dest
	if _, err := io.Copy(dfile, sfile); err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	// update dest permissions
	if err := dfile.Chmod(RwRR); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}
	return nil
}
