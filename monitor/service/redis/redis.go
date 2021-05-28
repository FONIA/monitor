package redis

import (
	. "fmt"
	"strings"
	"strconv"
	"time"
	cfg "monitor/config"
	"monitor/tool"
	"github.com/go-redis/redis"
)

//运行
func Run() {
	//运行监测
	c := cfg.RedisCheck{}
	c.Redisinit()
	checkredis(c)
}

//执行检测
func checkredis(c cfg.RedisCheck) {
	d:=time.Duration(c.CheckTime)*time.Second
	Println("[success] Redis检测服务已启动")
	tool.Log("info", "已运行监控项：Redis")
	// 连接redis
	dbcfg := strings.Fields(c.Url)
	if dbcfg[1] == "null" {
		dbcfg[1] = ""
	}
	db,_ := strconv.Atoi(dbcfg[2])
	for  {
		client := redis.NewClient(&redis.Options{
			Addr:     dbcfg[0],
			Password: dbcfg[1], // no password set
			DB:       db,  // use default DB
		})	
		_, err := client.Ping().Result()

		if err != nil {
			tool.Log("redis_err","redis检测故障:"+err.Error())
			if c.IsFailedReload{
					killredis(c)
					startredis(c)
				tool.Log("info","已重新启动redis")
			}

			if c.IsSendMsg && c.Msgtype == "email" {
				Notice := cfg.EmailNotice{}
				Notice.EmailInit()
				Notice.Title = "Redis服务掉线";
				Notice.Content = "ERR："+err.Error();
				tool.SendEmail(Notice)
			}
		}
		client.Close()
		time.Sleep(d)
	}

}

//杀掉服务
func killredis(c cfg.RedisCheck) {
	n, _ := tool.Kill(c.ServerName)
	str := "重启停止服务redis" + n
	tool.Log("info",str)

}

//启动
func startredis(c cfg.RedisCheck) {
	n,_:=tool.RunCommand(c.RestartShell)
	str:="启动redis服务"+n
	tool.Log("info",str)
}
