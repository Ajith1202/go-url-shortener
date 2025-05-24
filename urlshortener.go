package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	urlInput, _ := reader.ReadString('\n')
	urlInput = strings.TrimSpace(urlInput)

	urlMap := make(map[string]string)

	if isValidUrl(urlInput) {
		random_bytes := make([]byte, 32)

		_, err := rand.Read(random_bytes)
		if err != nil {
			panic(err)
		}

		shortUrl := base64.URLEncoding.EncodeToString(random_bytes)[:7]
		urlMap[shortUrl] = urlInput

		fmt.Println(shortUrl + " -> " + urlInput)
	}

}

func isValidUrl(urlInput string) bool {
	url, err := url.ParseRequestURI(urlInput)
	return err == nil && url.Scheme != "" && url.Host != ""
}
