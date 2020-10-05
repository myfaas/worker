package sandbox

import (
	"log"
	"strings"
	"../config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/mount"
	"golang.org/x/net/context"
)

type UpContainer struct {
	containerID string
	dependencies []string
}

var UpContainersPool = make([]UpContainer, 0, config.MaxUpContainersNum)

func InitPool(cli *client.Client) {
	
	ctx := context.Background()

	log.Println("Initializing a pool of docker containers...")
	
	// 创建所有容器 
	for _, image_name := range config.ImagesName {
		// 创建容器
		resp, err := cli.ContainerCreate(
			ctx, 
			&container.Config {
				Image: image_name,
				Cmd:   []string{"/bin/bash"},
				Tty:   true,
			}, 
			&container.HostConfig {
				Mounts: []mount.Mount {
					{
						Type:   mount.TypeBind,
						Source: config.LambdasDirPath,
						Target: "/tmp/",
					},
				},
			}, nil, "")

		if err != nil {
			panic(err)
		}

		// 启动容器
		if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

		// 把容器信息添加到运行容器池中
		log.Println(resp.ID)
		// 这里的初始容器都最多只有一个依赖
		deps := strings.Split(image_name, "-")[1:]
		log.Println(deps)
		up_container := UpContainer {resp.ID, deps}
		UpContainersPool = append(UpContainersPool, up_container)
	}

	log.Println("Initialize Done...")

}
