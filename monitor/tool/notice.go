package tool

import (
	"monitor/config"
	"net/smtp"
	"time"
	"fmt"
)

//发送微信消息  url地址  msg 消息内容
func SendWx(msg string, phone ...string) {
	url := config.Wxurl
	wphone:=config.WxPhone
	maps := make(map[string]interface{})
	maps["msgtype"] = "text"
	maps2 := make(map[string]interface{})
	maps2["content"] = msg
	phonelist := []string{}
	for _, v := range phone {
		phonelist = append(phonelist, v)
	}
	phonelist=append(phonelist,wphone)
	maps2["mentioned_mobile_list"] = phonelist
	maps["text"] = maps2
	res :=Post(url, maps, "application/json")
	Log("info", res)
}

//发email
func SendEmail(data config.EmailNotice, mail ...string) {
	// 控制频率0-15、35-50
	t := time.Now().Second()
	if (t > 15 && t < 35) || (t > 50) {
		return
	}
	Log("send_mail", "开始发送报警邮件")
	// 认证, content-type设置
	auth := smtp.PlainAuth("", data.MailUser, data.MailPwd, data.MailHost)
	contentType := "Content-Type: text/html; charset=UTF-8"
	list := []string{}
	for _, v := range mail {
		list = append(list, v)
	}
	list=append(list,data.Target)

	ip,_ := getLocalIPv4Address()
	title := data.Title + " IP：" + ip
	for _, to := range list {
		s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s", to, data.MailNick, data.MailUser, title, contentType, data.Content)
		msg := []byte(s)
		addr := fmt.Sprintf("%s:%s", data.MailHost, data.MailPort)
		err := smtp.SendMail(addr, auth, data.MailUser, []string{to}, msg)
		if err != nil {
			Log("send_mail_err", err.Error())
		}
	}
}
