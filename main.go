package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"workspace/cloud-pan/config"
	_ "workspace/cloud-pan/db/mysql"
	"workspace/cloud-pan/router"
	"workspace/cloud-pan/util/logs"
)

func main() {
	err := Run()
	if err != nil {
		log.Print(err)
	}
}

func Run() (err error) {
	//初始化oss

	//初始化mysql

	//初始化redis

	//初始化log
	logs.Init("./log", "log", 3, false)
	//捕捉退出信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	// 启动服务监听
	go func(port int, mode string) {
		if err := router.RouterRun(port, mode); err != nil {
			log.Printf("server run failed:%v", err)
			return
		}
	}(config.HTTP_PORT, "dev")

	log.Printf("server[%s] serving on http_port:%d", "dev", config.HTTP_PORT)
	s := <-sig
	log.Printf("receive signal:%d , server stopping ...", s)
	return nil
}

