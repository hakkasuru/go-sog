package core

type Core interface {
	Write(url string, title string, msg string, tags []string) error
}

type noopCore struct{}

func NewNoopCore() Core                                                            { return noopCore{} }
func (c noopCore) Write(url string, title string, msg string, tags []string) error { return nil }
