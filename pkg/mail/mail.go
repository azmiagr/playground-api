package mail

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

func SendEmail(to, subject, message string) error {
	SMTP_HOST := os.Getenv("SMTP_HOST")
	SMTP_PORT := os.Getenv("SMTP_PORT")
	SMTP_USERNAME := os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")

	addr := fmt.Sprintf("%s:%s", SMTP_HOST, SMTP_PORT)
	msg := fmt.Sprintf(
		"From: No Reply <%s>\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
			"\r\n%s", // body setelah header
		SMTP_USERNAME, to, subject, message)
	err := smtp.SendMail(addr,
		smtp.PlainAuth("", SMTP_USERNAME, SMTP_PASSWORD, SMTP_HOST),
		SMTP_USERNAME, []string{to}, []byte(msg))

	if err != nil {
		return err
	}

	return nil
}

func GenerateCode() string {
	minRange, maxRange := 100000, 999999

	return strconv.Itoa(rand.Intn(maxRange-minRange+1) + minRange)
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}

	return string(result)
}
