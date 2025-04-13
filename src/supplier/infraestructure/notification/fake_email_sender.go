package notification

import "fmt"

type FakeEmailSender struct{}

func (f *FakeEmailSender) SendEmail(to, subject, body string) error {
    fmt.Printf("[Simular envío de email]\nTo: %s\nSubject: %s\nBody: %s\n", to, subject, body)
    return nil
}
