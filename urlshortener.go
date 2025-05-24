package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"

	"github.com/gin-gonic/gin"
)

var urlMap = make(map[string]string)

func main() {

	restEngine := gin.Default()
	restEngine.POST("/urlShorten", func(ctx *gin.Context) {
		var req struct {
			URL string `json:"url"`
		}

		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(400, gin.H{"Error": "Invalid Request"})
			return
		}

		if !isValidUrl(req.URL) {
			ctx.JSON(400, gin.H{"Error": "Invalid Url"})
			return
		}

		shortUrl := getShortenedUrl()
		urlMap[shortUrl] = req.URL
		ctx.JSON(200, gin.H{"Shortened URL": "http://localhost:8080/" + shortUrl})
	})

	restEngine.GET("/:code", func(ctx *gin.Context) {
		code := ctx.Param("code")

		longUrl, exists := urlMap[code]
		if !exists {
			ctx.JSON(404, gin.H{"Error": "URL Not Found"})
			return
		}

		ctx.Redirect(302, longUrl)
	})

	restEngine.Run(":8080")
}

func getShortenedUrl() string {
	random_bytes := make([]byte, 32)

	_, err := rand.Read(random_bytes)
	if err != nil {
		panic(err)
	}

	shortUrl := base64.URLEncoding.EncodeToString(random_bytes)[:7]

	return shortUrl
}

func isValidUrl(urlInput string) bool {
	url, err := url.ParseRequestURI(urlInput)
	return err == nil && url.Scheme != "" && url.Host != ""
}
