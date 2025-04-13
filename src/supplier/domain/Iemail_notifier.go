package domain

type EmailNotifier interface{
	SendEmail(to, subject, body string) error
}