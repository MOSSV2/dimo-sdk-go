package build

var CurrentCommit string

const BuildVersion = "0.3.2-alpha"

func UserVersion() string {
	return BuildVersion + CurrentCommit
}
