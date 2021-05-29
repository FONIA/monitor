package server

import (
	. "fmt"
	"time"
	"monitor/config"
	"monitor/tool"
	"strconv"
)

func Run() {
	c := config.ServerCheck{}
	c.Serverinit() //初始化
	checkservice(c) //执行检测
}

//检测服务
func checkservice(data config.ServerCheck) {
	d := time.Duration(data.CheckTime) * time.Second
	Println("[success] 服务器检测服务已启动")
	tool.Log("info", "已运行监控项：Server")
	for {
		cpu := tool.GetCpuPercent() //检测CPU
		mem := tool.GetMemPercent() //检测内存
		disk := tool.GetDiskPercent() //检测硬盘
		if cpu >= data.Cpu || mem >= data.Mem || disk >= data.Disk {
			if data.IsSendMsg && data.Msgtype == "email" {
				Notice := config.EmailNotice{}
				Notice.EmailInit()
				Notice.Title = "服务器超载";
				Notice.Content = "CPU利用率：" + strconv.FormatFloat(cpu, 'E', -1, 64) + " 内存利用率：" + strconv.FormatFloat(mem, 'E', -1, 64) + " 磁盘利用率：" + strconv.FormatFloat(disk, 'E', -1, 64);
				tool.SendEmail(Notice)
			}
		}
		time.Sleep(d)
	}

}
