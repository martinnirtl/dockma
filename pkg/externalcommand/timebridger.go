package externalcommand

// Timebridger describes interface for time bridging implementations
type Timebridger interface {
	Start(command string) error
	Update(update string) error
	Stop() error
}
