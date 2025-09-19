package version

import (
	"fmt"
	"runtime"
)

const unknown = "unknown"

var (
	Version   = unknown
	BuildTime = unknown
	GitCommit = unknown
	GitBranch = unknown
	GoVersion = runtime.Version
)

type AppInfo struct {
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
	GitCommit string `json:"got_commit"`
	GitBranch string `json:"git_branch"`
	GoVersion string `json:"go_version"`
}

func Get() AppInfo {
	return AppInfo{
		Version:   Version,
		BuildTime: BuildTime,
		GitCommit: GitCommit,
		GitBranch: GitBranch,
		GoVersion: GoVersion(),
	}
}

func (i AppInfo) String() string {
	return fmt.Sprintf("Version: %s, Build: %s, Commit: %s, Branch: %s, Go: %s", i.Version, i.BuildTime, i.GitCommit, i.GitBranch, i.GoVersion)
}
