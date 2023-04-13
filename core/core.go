package core

// Core logging interface
type Core interface {
	// Write message into payload and publish to slack webhook
	Write(url string, title string, msg string, tags []string) error
}

type noopCore struct{}

// NewNoopCore no operation core
func NewNoopCore() Core                                                            { return noopCore{} }
func (c noopCore) Write(url string, title string, msg string, tags []string) error { return nil }
