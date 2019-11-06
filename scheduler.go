package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
)

func myProxy() gin.HandlerFunc {

    target := "192.168.74.129:8080"

    return func(c *gin.Context) {
        director := func(req *http.Request) {
            r := c.Request
            req.URL.Scheme = "http"
            req.URL.Host = target
            req.Header["my-header"] = []string{r.Header.Get("my-header")}
                        // Golang camelcases headers
            delete(req.Header, "My-Header")
		}

		fmt.Println("aha, i got you!")
		fmt.Println(c.Param("name"))

        proxy := &httputil.ReverseProxy{Director: director}
        proxy.ServeHTTP(c.Writer, c.Request)
    }
}

func main() {

	router := gin.Default()
	router.GET("/lambda/:name", myProxy())
	router.Run() // listen and serve on 0.0.0.0:8080
}
