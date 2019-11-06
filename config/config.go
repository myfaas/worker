package config

const PackageNumber int = 7
const LambdasDirPath string = "/home/wen/myfaas/lambdas/"

// 镜像都是提前在系统中准备好的
var ImagesName = [...]string {"ubuntu:python3-pip", "ubuntu:python3-simplejson"}
	// "ubuntu:python3-watchdog",
	// "ubuntu:python3-simplejson",
	// "ubuntu:python3-arrow",


const MaxUpContainersNum int = 15

