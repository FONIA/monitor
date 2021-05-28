package mysql

import (
	"database/sql"
	. "fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"monitor/config"
	"monitor/tool"
)

//运行
func Run() {
	//运行监测
	c := config.MysqlCheck{}
	c.Mysqlinit()
	checkmysql(c)
}

//执行检测
func checkmysql(c config.MysqlCheck) {
	d:=time.Duration(c.CheckTime)*time.Second
	Println("[success] Mysql检测服务已启动")
	tool.Log("info", "已运行监控项：Mysql")
	for  {
		db, _ := sql.Open("mysql", c.Url)
		err := db.Ping()
		if err != nil {
			tool.Log("mysql_err","mysql检测故障:"+err.Error())
			if c.IsFailedReload{
					killmysql(c)
					startmysql(c)
				tool.Log("info","已重新启动mysql")
			}

			if c.IsSendMsg && c.Msgtype == "email" {
				Notice := config.EmailNotice{}
				Notice.EmailInit()
				Notice.Title = "Mysql服务掉线";
				Notice.Content = "ERR："+err.Error();
				tool.SendEmail(Notice)
			}
		}
		db.Close()
		time.Sleep(d)
	}

}

//杀掉服务
func killmysql(c config.MysqlCheck) {
	n, _ := tool.Kill(c.ServerName)
	str := "重启停止服务mysql" + n
	tool.Log("info",str)

}

//启动mysql
func startmysql(c config.MysqlCheck) {
	n,_:=tool.RunCommand(c.RestartShell)
	str:="启动mysql服务"+n
	tool.Log("info",str)

}
