服务监控（php+mysql+nginx+redis+服务器超载）
====



1、特性
----


- 1、目前支持监控mysq、nginx、php 、redis、cpu、cache、disk
- 2、自动重启功能 监测服务不可用自动重启
- 3、新增消息通知功能 支持企业微信消息推送、邮件推送（qq邮箱用587端口）



2、安装方法
----

获取安装

- go get github.com/FONIA/monitor
- go get github.com/go-redis/redis
- go get github.com/shirou/gopsutil
- go get github.com/go-sql-driver/mysql
- 修改config 配置文件
- 赋权：chmod +x run.sh stop.sh
- 启动服务 sh run.sh 服务名 
- 后台运行 nohup sh run.sh 服务名 &
- 停止服务 sh stop.sh
  
