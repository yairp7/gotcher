package version

// Version information
var (
	// Version is the current version of Gotcher
	Version = "0.1.0"

	// GitCommit is the git commit hash and will be filled in by the build system
	GitCommit = "unknown"

	// BuildDate is the date when the binary was built and will be filled in by the build system
	BuildDate = "unknown"
)

// GetVersionInfo returns a formatted string containing version information
func GetVersionInfo() string {
	return "Gotcher version " + Version + " (commit: " + GitCommit + ", built: " + BuildDate + ")"
}
