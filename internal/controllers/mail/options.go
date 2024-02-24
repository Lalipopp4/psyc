package mail

import "net/smtp"

func (m *mailManager) Send(receiver string, data []byte) error {
	return smtp.SendMail(m.addr, m.auth, m.mail, []string{receiver}, data)
}
