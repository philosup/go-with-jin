# go-with-jin
example

## Go 언어 설치 https://golang.org/doc/install

download go for windows - 1.15.6

## Visual Studio Code
- Go for Visual Studio Code

https://github.com/golang/vscode-go 참조

> Ctrl+Shift+P

> Go: Install/Update Tools 

> all select -> 17개

## 설치
https://github.com/gin-gonic/gin

Installation
To install Gin package, you need to install Go and set your Go workspace first.

The first need Go installed (version 1.12+ is required), then you can use the below Go command to install Gin.
$ go get -u github.com/gin-gonic/gin
Import it in your code:
import "github.com/gin-gonic/gin"
(Optional) Import net/http. This is required for example if using constants such as http.StatusOK.
import "net/http"