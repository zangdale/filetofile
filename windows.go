//go:build windows
// +build windows

package main

func getFilePath(dir, arg string) (string, error) {
	return filepath.Join(dir, arg), nil
}
