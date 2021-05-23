package nginx

import (
	. "fmt"
	"time"
	"monitor/config"
	"monitor/tool"
)

func Run() {
	c := config.NginxCheck{}
	c.Nginxinit() //初始化
	checkservice(c) //执行检测
}

//检测服务
func checkservice(data config.NginxCheck) {
	d := time.Duration(data.CheckTime) * time.Second
	Println("[success] Nginx检测服务已启动")
	tool.Log("info", "已运行监控项：Nginx")
	for {
		status := tool.Get(data.Url)
		if !status {
			if data.IsSendMsg && data.Msgtype == "email" {
				Notice := config.EmailNotice{}
				Notice.EmailInit()
				Notice.Title = "Nginx服务掉线";
				Notice.Content = "宕机URL："+data.Url;
				tool.SendEmail(Notice)
			}
			if data.IsFailedReload{
				reloadshll(data)   //杀掉服务
				startservice(data) //启动服务
			}
		}
		time.Sleep(d)
	}

}

//停止服务
func reloadshll(data config.NginxCheck)  {
	n, _ := tool.Kill(data.ServerName)
	stres:="停止服务nginx..."+ n
	tool.Log("info",stres)

}

//启动服务
func startservice(data config.NginxCheck) {
	n,_:=tool.RunCommand(data.RestartShell)
	stre:="启动服务nginx..."+ n
	tool.Log("info",stre)
}
