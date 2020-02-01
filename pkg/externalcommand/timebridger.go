package externalcommand

import "io"

// Timebridger describes interface for time bridging implementations
type Timebridger interface {
	Start(description string) (Output, error)
	Update(description string) error
	Stop() error
}

// Output is used for rewriting stdout and stderr of external command by a time bridger implementation
type Output struct {
	out io.Writer
	err io.Writer
}
