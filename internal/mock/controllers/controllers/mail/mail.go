package mail

type MailSender interface {
	Send(receiver string, data []byte) error
}
