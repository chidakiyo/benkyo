package main

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"os"
)

func main() {

	domain := os.Getenv("MAILGUN_DOMAIN")
	key := os.Getenv("MAILGUN_KEY")
	to := os.Getenv("MAIL_TO")

	fmt.Printf("%s, %s, %s\n", domain, key, to)

	mg := mailgun.NewMailgun(
		domain,
		key,
	)

	message := mg.NewMessage(
		"test@" + domain,
		"メールテスト",
		"テストメールを送信します\ntest.",
		to,
	)

	message.AddHeader("Return-Path", "test@example.com")

	mes, id, err := mg.Send(message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", mes)
	fmt.Printf("%v", id)
}
