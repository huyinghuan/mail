## 邮件发送服务

不建议使用此库，推荐:https://github.com/go-gomail/gomail

适用于国内的 163， qq邮箱

不支持附件

### 使用

```
import "github.com/huyinghuan/mail"

func TestSender(t *testing.T) {
	sender := &mail.MailSender{
		Username: "xxxxx@163.com",
		Password: "xx",
		Host:     "smtp.163.com",
		Port:     465}
	/*
		sender := &mailMailSender{
			Username: "xxxxx@qq.com",
			Password: "xx",
			Host:     "smtp.qq.com",
			Port:     465}
	*/
	if err := sender.SendMail("GOMAIL Test", "hello world!", []string{"ec.huyinghuan@gmail.com"}); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		log.Println("发送成功！")
	}
}

如果抛出 unencrypted connection 错误，请设置  `sender.NoSSL = true`
```

## LICENSE

MIT
