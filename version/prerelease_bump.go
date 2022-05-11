package version

import "github.com/blang/semver"

type PreReleaseBump struct {
	Pre string
}

func (bump PreReleaseBump) Apply(v semver.Version) semver.Version {
	if v.Pre == nil {
		v.Patch++
	}

	if v.Pre == nil || v.Pre[0].VersionStr != bump.Pre {
		v.Pre = []semver.PRVersion{
			{VersionStr: bump.Pre},
			{VersionNum: 1, IsNum: true},
		}
	} else {
		if len(v.Pre) < 2 {
			v.Pre = append(v.Pre, semver.PRVersion{VersionNum: 0, IsNum: true})
		}

		v.Pre = []semver.PRVersion{
			{VersionStr: bump.Pre},
			{VersionNum: v.Pre[1].VersionNum + 1, IsNum: true},
		}
	}

	return v
}
