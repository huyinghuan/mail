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

	if err := sender.SendMail("GOMAIL Test", "hello world!", []string{"ec.huyinghuan@gmail.com", "646344359@qq.com"}); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		log.Println("发送成功！")
	}
}
