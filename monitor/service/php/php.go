package php

import (
	. "fmt"
	"time"
	"monitor/config"
	"monitor/tool"
)

func Run()  {
	c:=config.PhpCheck{}
	c.Phpinit()
	phpcheck(c)
}

func phpcheck(c config.PhpCheck)  {
	d:=time.Duration(c.CheckTime)*time.Second
	Println("[success] php检测服务已启动")
	tool.Log("info", "已运行监控项：php")
	for{
		code,status:=tool.GetData(c.Url)
		if code==0{
			time.Sleep(d)
			return
		}
		if code!=200{
			tool.Log("php_err","php检测故障:"+status)
			if c.IsFailedReload{
				tool.Kill(c.ServerName)
				tool.RunCommand(c.RestartShell)
				tool.Log("info","重新启动php")
			}

			if c.IsSendMsg && c.Msgtype == "email" {
				Notice := config.EmailNotice{}
				Notice.EmailInit()
				Notice.Title = "php服务掉线";
				Notice.Content = "ERR："+status;
				tool.SendEmail(Notice)
			}
		}

		time.Sleep(d)
	}

}