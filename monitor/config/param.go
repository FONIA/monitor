package config

type Configs struct {
	ServerName   string  //进程名称(需要与linux 进程名相同否则kill重启进程false)
	Url            string //检测地址
	Msgtype		   string // 通知方式
	IsSendMsg      bool   //服务异常是否开启通知
	IsFailedReload bool   //服务异常是否重新启动
	RestartShell   string //服务重启命令
	CheckTime int64   //检测间隔时间
}

type Email struct {
	Title	string //邮件标题
	Content string //邮件内容
	MailHost string //邮件服务器地址
	MailPort string //端口
	MailUser string //发件邮箱账号
	MailPwd  string //发件邮箱密码
	MailNick string //发件昵称
	Target string //收件人
}