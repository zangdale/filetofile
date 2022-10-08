package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("input file")
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir, args[1])
	path, err := getFilePath(dir, args[1])
	if err != nil {
		log.Fatal(err)
	}
	bodyBase64, err := getFileBodyToBase64(path)
	if err != nil {
		log.Fatal(err)
	}

	paths, fileName := filepath.Split(path)
	// fmt.Println(paths, fileName)
	goFile := filepath.Join(paths, strings.ReplaceAll(fileName, ".", "_")+".go")
	body := `package main

import "encoding/base64"

var FileBodyBase64Str = "` + bodyBase64 + `"

var FileBodyByte, _ = base64.StdEncoding.DecodeString(FileBodyBase64)

`
	err = ioutil.WriteFile(goFile, []byte(body), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(goFile)
}

func getFileBodyToBase64(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
