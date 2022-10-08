# filetofile

## install
```
go install github.com/zangdale/filetofile
```

## build
```
go build .
```

## use

```
go run . image.jpeg
or 
filetofile image.jpeg
```
export file 
```
package main

import "encoding/base64"

var FileBodyBase64Str = "/9j/***********jIyMjIyMjIyMjIy"

var FileBodyByte, _ = base64.StdEncoding.DecodeString(FileBodyBase64)

```