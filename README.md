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

## Gin 설치
https://github.com/gin-gonic/gin

The first need Go installed (version 1.12+ is required), then you can use the below Go command to install Gin.

> go get -u github.com/gin-gonic/gin


main.go
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

```
> go run main.go