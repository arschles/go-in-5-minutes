package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func kittyHandler(c *gin.Context) {
	cats := []string{
		"/img/kitties/1.jpg",
		"/img/kitties/2.jpg",
		"/img/kitties/3.jpg",
		"/img/kitties/4.jpg",
		"/img/kitties/5.jpg",
		"/img/kitties/6.jpg",
	}
	htmlStr := fmt.Sprintf(`<html>
	<head></head>
	<body>
	<center>
	<img src="%s" width="1000"/>
	</center>
	</body>
	</html>`, randStr(cats))
	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(htmlStr))
}
