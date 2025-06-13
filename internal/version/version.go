package version

import "runtime/debug"

// Version information
var (
	// Version is the current version of Gotcher
	Version = "dev"

	// GitCommit is the git commit hash and will be filled in by the build system
	GitCommit = "unknown"

	// BuildDate is the date when the binary was built and will be filled in by the build system
	BuildDate = "unknown"
)

func GetVersion() string {
	if Version != "dev" {
		return Version
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}

	return "unknown"
}
