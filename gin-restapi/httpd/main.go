package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/go-samples/gin-restapi/article"
	"github.com/tanimutomo/go-samples/gin-restapi/httpd/handler"
)

func main() {
	article := article.New()
	r := gin.Default()
	r.GET("/article", handler.ArticleGet(article))
	r.POST("/article", handler.ArticlePost(article))

	r.Run()
}
