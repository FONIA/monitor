package config

//企业微信机器人消息推送地址申请文档https://work.weixin.qq.com/api/doc/90000/90136/91770
var Wxurl="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=d4f814"
var WxPhone="xxx" //@提醒人 可为空
var Service = map[string]bool{
	"nginx":true,
	"mysql":true,
	"php":true,
}

//nginx配置监控参数
type NginxCheck struct {
	Configs
}
func (c *NginxCheck) Nginxinit() {
	c.ServerName="nginx"  //服务名称
	c.Msgtype="email"     // 默认email
	c.IsSendMsg=true	 // 邮件
	c.Url = "http://127.0.0.1:82" //检测地址 
	c.IsFailedReload = true ////服务异常是否重新启动
	c.RestartShell = "docker restart nginx"  //重启服务命令
	c.CheckTime=5  //检测间隔时间（秒）
}

//邮箱报警
type EmailNotice struct {
	Email
}

func (c *EmailNotice) EmailInit() {
	c.MailHost = "smtp.qq.com"
	c.MailPort = "587"
	c.MailUser = "1033569557@qq.com"
	c.MailPwd = "xxx"
	c.MailNick = "FONIA"
	c.Target = "1033569557@qq.com"
}

//mysql检测
type MysqlCheck struct {
	Configs
}

func (m *MysqlCheck) Mysqlinit()  {
	m.ServerName="mysql"  //服务名称
	m.Msgtype="email"     // 默认email
	m.Url = "root:root(127.0.0.1:3306)/godata" //mysql账号密码数地址端口数据库
	m.IsFailedReload = true ////服务异常是否重新启动
	m.IsSendMsg=true
	m.RestartShell = "service mysqld restart"  //重启mysql服务命令
	m.CheckTime=5  //检测间隔时间（秒）
}

//php检测
type PhpCheck struct {
	Configs
}

func (m *PhpCheck) Phpinit()  {
	m.ServerName="php"  //服务名称
	m.Msgtype="email"     // 默认email
	m.Url = "http://127.0.0.1/a.php" //地址检测
	m.IsFailedReload = true ////服务异常是否重新启动
	m.IsSendMsg=true
	m.RestartShell = " /usr/local/php/sbin/php-fpm && /usr/local/php5.6/sbin/php-fpm "  //重启服务命令
	m.CheckTime=5  //检测间隔时间（秒）
}
