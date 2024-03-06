package versioning

var (
	// Version is the main version at the moment.
	// Commit is the git commit that the binary was built on
	// BuildTime is the timestamp of the build
	// Embedded by --ldflags on build time
	// Versioning should follow the SemVer guidelines
	// https://semver.org/
	Version   = "1.0.4_beta"
	Branch    = "dev"
	BuildTime = "1/12/2023"
	Commit    string
)
