package main

import (
	"github.com/gin-gonic/gin"
	"github.com/docker/docker/client"
	"net/http"
	"./sandbox"
)

func main() {

	// 启动一个docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// 初始化, 开启一个容器池
	sandbox.InitPool(cli)

	// 监听请求
	r := gin.Default()
	r.GET("/lambda/:name", func(c *gin.Context) {
		lambda_name := c.Param("name")
		res := sandbox.ExecLambda(cli, lambda_name)
		c.String(http.StatusOK, res)
	})
	
	// 默认在8080端口
	r.Run()
	
}
