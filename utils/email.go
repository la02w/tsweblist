package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

const text = `<p>Teamspeak连接地址：<span style="color: #86b300;">%s</span></p><p>复制到浏览器打开</p>`

func SeedEmail(email string, url string) {
	m := gomail.NewMessage()
	m.SetHeader("From", SMTPEMAIL)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "TeamSpeakList")
	m.SetBody("text/html", fmt.Sprintf(text, url))
	//附件
	//m.Attach("./myIpPic.png")
	d := gomail.NewDialer(SMTPHOST, SMTPPORT, SMTPEMAIL, SMTPPASSWORD)
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
}
