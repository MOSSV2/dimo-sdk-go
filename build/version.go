package build

var CurrentCommit string

const BuildVersion = "0.4.0-alpha"

func UserVersion() string {
	return BuildVersion + CurrentCommit
}
