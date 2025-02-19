package build

var CurrentCommit string

const BuildVersion = "0.4.1-alpha"

func UserVersion() string {
	return BuildVersion + CurrentCommit
}
