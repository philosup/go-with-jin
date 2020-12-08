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
    
    r.GET("/time", api.Time)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

```
> go run main.go

http://localhost:8080/ping

## MongoDB  

### 설치 https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/

다운로드 
https://www.mongodb.com/try/download/community?tck=docs_server

MongoDB Community Server 4.4.2

https://m.blog.naver.com/wideeyed/221815886721



https://www.mongodb.com/golang

> go get go.mongodb.org/mongo-driver