package version

import "github.com/blang/semver"

type FinalBump struct{}

func (FinalBump) Apply(v semver.Version) semver.Version {
	if v.Pre == nil {
		v.Patch++
	} else {
		v.Pre = nil
	}

	return v
}
