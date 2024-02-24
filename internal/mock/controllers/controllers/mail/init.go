package mail

type mailManager struct {
}

func New(cfg *Config) *mailManager {
	return &mailManager{}
}
