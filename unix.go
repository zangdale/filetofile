//go:build linux || darwin
// +build linux darwin

package main

import (
	"path/filepath"
)

func getFilePath(dir, arg string) (string, error) {
	// fmt.Println(filepath.Abs(arg))
	return filepath.Join(dir, arg), nil
}
