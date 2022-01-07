package config

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Release vars - injected by the compiler
var (
	CommitHash = ""
	ReleaseCid = ""
)

// ReleaseSummary contains concise release info.
type ReleaseSummary struct {
	Commit string `json:"commit,omitempty"`
	IPFS   string `json:"ipfs,omitempty"`
	// Version string `json:"version,omitempty"` TODO: Use this when semver is injected
}

func GetBuildReleaseInfo() (*ReleaseSummary, bool) {
	if len(CommitHash) == 0 {
		return nil, false
	}

	return &ReleaseSummary{
		Commit: CommitHash,
		IPFS:   ReleaseCid,
	}, true
}

// ReleaseInfo contains the release response from the updater.
type ReleaseInfo struct {
	IPFS     string          `json:"ipfs"`
	Manifest ReleaseManifest `json:"manifest"`
}

// String implements fmt.Stringer interface.
func (releaseInfo *ReleaseInfo) String() string {
	if releaseInfo == nil {
		return ""
	}
	b, _ := json.Marshal(releaseInfo)
	return string(b)
}

// ReleaseInfoFromString parses the string.
func ReleaseInfoFromString(s string) *ReleaseInfo {
	if len(s) == 0 {
		log.Warn("empty release info")
		return nil
	}
	var releaseInfo ReleaseInfo
	json.Unmarshal([]byte(s), &releaseInfo)
	if len(releaseInfo.Manifest.Release.Commit) > 0 {
		log.WithFields(log.Fields{
			"commit": releaseInfo.Manifest.Release.Commit,
			"ipfs":   releaseInfo.IPFS,
		}).Info("found release info")
	}
	return &releaseInfo
}

// MakeSummaryFromReleaseInfo transforms the release info into a more compact and common form.
func MakeSummaryFromReleaseInfo(releaseInfo *ReleaseInfo) *ReleaseSummary {
	if releaseInfo == nil {
		return nil
	}
	return &ReleaseSummary{
		Commit: releaseInfo.Manifest.Release.Commit,
		IPFS:   releaseInfo.IPFS,
	}
}

// ReleaseManifest contains the latest info about the latest scanner version.
type ReleaseManifest struct {
	Release Release `json:"release"`
}

// Release contains release data.
type Release struct {
	Timestamp  string          `json:"timestamp"`
	Repository string          `json:"repository"`
	Commit     string          `json:"commit"`
	Services   ReleaseServices `json:"services"`
}

// ReleaseServices are the services to run for scanner node.
type ReleaseServices struct {
	Updater    string `json:"updater"`
	Supervisor string `json:"supervisor"`
}