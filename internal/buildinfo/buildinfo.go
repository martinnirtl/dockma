package buildinfo

import (
	"fmt"
	"time"
)

// VersionInfo describes structure of version information objects.
type VersionInfo struct {
	Version string
	Commit  string
	Date    string
}

var (
	version = "development"
	commit  = "na"
	date    = time.Now().Local().Format(time.ANSIC)

	// Version provides access to version info and related printing functions.
	Version *VersionInfo
)

func init() {
	Version = &VersionInfo{
		Version: version,
		Commit:  commit,
		Date:    date,
	}
}

// Print prints the version information.
func (v *VersionInfo) Print() {
	fmt.Printf("Dockma CLI version %s\n", v.Version)
}

// PrintFull prints the full version information including built date and last commit.
func (v *VersionInfo) PrintFull() {
	fmt.Printf("Dockma CLI Version %s built on %s (commit: %s)\n", v.Version, v.Date, v.Commit)
}
