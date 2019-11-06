package sandbox

import (
	"os"
	"log"
	"bufio"
	"strings"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"../config"
)


/*
** <声明>
** 函数依赖：通过拉取代码文件，然后解析代码的第一行获得
** 选择容器：工作节点内部维护一份依赖包和容器ID的数据结构（因为每个工作节点的容器总数比较少，这个查找很快）
*/
func ExecLambda(cli *client.Client, name string) string {
	/* 根据要执行的函数名称，获取执行沙箱，然后执行代码，返回结果 */
	
	/* 第1步：拉取代码，解析函数依赖集合*/
	// 先判断是否已经存在
	filepath := config.LambdasDirPath + name + ".py"
	depends := make([]string, 0, config.PackageNumber)
	if _, err := os.Stat(filepath); err == nil {// 文件已经存在了
		depends = readPyFile(filepath)
	} else if os.IsNotExist(err) {// 文件不存在，需要从代码仓库拉取
		// TODO
		log.Println("No")
	} else {
		log.Println(err)
		return "Error. Failed to read Python file."
	}
	log.Println("------------")
	log.Println(depends)

	// TODO 第2步：获取要执行的沙箱
	containerID := UpContainersPool[1].containerID
	// 第3步：执行代码，获得结果，返回
	return execPython(cli, name, containerID)
}


// 根据python文件路径，读取依赖列表
// 依赖包需要按照规定格式，# dependency: dep1, dep2，注意使用逗号和空格一起隔开
func readPyFile(filepath string) (deps_slice []string) {
	deps_slice = nil
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		first_line := scanner.Text()
		deps := first_line[len("# dependency: "):]
		deps_slice = strings.Split(deps, ", ")
		// break
		return
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}


func execPython(cli *client.Client, name, containerID string) string {
	// 创建一个exec process
	exec_config := types.ExecConfig {AttachStdout: true, Tty: true, Cmd: []string {"python3", "/tmp/"+name+".py"}}
	idResponse, _ := cli.ContainerExecCreate(context.Background(), containerID, exec_config)
	// 执行exec，获取输出结果
	hijackedResponse, _ := cli.ContainerExecAttach(context.Background(), idResponse.ID, types.ExecStartCheck{false, true})
	res := make([]byte, 64)
	hijackedResponse.Reader.Read(res)
	result := string(res[:len(name)+5])
	return result
}
