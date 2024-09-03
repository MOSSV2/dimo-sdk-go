package build

var CurrentCommit string

const BuildVersion = "0.3.2"

func UserVersion() string {
	return BuildVersion + CurrentCommit
}
