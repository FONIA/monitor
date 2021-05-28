package main

import (
	. "fmt"
	"os"
	"monitor/tool"
	"monitor/service/php"
	"monitor/service/nginx"
	"monitor/service/mysql"
	"monitor/service/redis"
	"monitor/service/server"
	cfg "monitor/config"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		Println("\r\n入参错误：\r\n1.服务监控项：", cfg.Service,  "\r\n2.终止脚本：stop")
		return
	}

	if arg[1] == "stop" {
		Println("服务停止...")
		stop()
		return
	}

	str := make(map[string]bool)
	for _,v := range arg{
		str[v] = true
	}

	for key := range str{
		_,ok := cfg.Service[key]
		if ok==true {
			if cfg.Service[key] == true {
				Registered(key)
			} else {
				Println("不支持的监控项：", key)
			}
		}
	}
	select {}
}

func stop() {
	_, err := tool.Kill("ServiceMonitorMain")
	if err != nil {
		pid,_ := tool.GetPid("ServiceMonitorMain")
		Println("停止失败，请手动Kill\r\n进程名：ServiceMonitorMain\r\n进程ID：", pid)
	}
}

func Registered(Service string) {
	switch Service {
		case "nginx":
			go nginx.Run() //注册nginx检测
		case "mysql":
			go mysql.Run() //注册mysql检测
		case "php":
			go php.Run()   //注册php检测
		case "redis":
			go redis.Run() //redis
		case "server":
			go server.Run()	//服务器cpu、cache、disk监控
		default:
			Println("不存在的监控项：", Service)
	}
}