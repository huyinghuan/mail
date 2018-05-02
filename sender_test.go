package mail

import "testing"
import "log"

func TestSender(t *testing.T) {
	sender := &MailSender{
		NoSSL: true,
		Username: "xiacijian@163.com",
		Password: "xxx",
		Host:     "smtp.163.com",
		Port:     465}

		//sender := &MailSender{
		//	Username: "646344359@qq.com",
		//	Password: "",
		//	Host:     "smtp.qq.com",
		//	Port:     465}

	if err := sender.SendMail("GOMAIL Test", "hello world!", []string{"ec.huyinghuan@gmail.com", "646344359@qq.com"}, map[string]string{
		"Content-Type":  "text/plain; charset=\"utf-8\"",
	}); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		log.Println("发送成功！")
	}
}

func TestSender2(t *testing.T) {
	sender := &MailSender{
		NoSSL: true,
		Username: "xiacijian@163.com",
		Password: "xxx",
		Host:     "smtp.163.com",
		Port:     465}

	newMail := MailContent{
		Subject:"GOMAIL Test",
		Content: "hello world!",
		Contact: []string{"ec.huyinghuan@gmail.com", "646344359@qq.com"},
		Other: map[string]string{
			"Content-Type":"text/plain; charset=\"utf-8\"",
		},
	}

	if err := sender.Send(&newMail); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		log.Println("发送成功！")
	}
}

