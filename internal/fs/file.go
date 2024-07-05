package fs

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copies a provided file from src to dest with a default permission of 0o644. It fails if it's a directory.
func CopyFile(src, dest string) error {
	// read file from fsys (OperatingFS or specific fsys)
	sfile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", src, err)
	}
	defer sfile.Close()

	// create dest in OS filesystem and not given fsys
	dfile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", dest, err)
	}
	defer dfile.Close()

	// copy buffer from src to dest
	if _, err := io.Copy(dfile, sfile); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	// update dest permissions
	if err := dfile.Chmod(RwRR); err != nil {
		return fmt.Errorf("failed to update %s permissions: %w", dest, err)
	}
	return nil
}

// Exists returns a boolean indicating whether the provided input src exists or not.
func Exists(src string) bool {
	// read file from fsys (OperatingFS or specific fsys)
	file, err := os.Open(src)
	if err != nil {
		return false
	}
	_ = file.Close()
	return true
}
