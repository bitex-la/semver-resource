package version

func BumpFromParams(bumpStr string, preStr string) Bump {
	var semverBump Bump

	switch bumpStr {
	case "final":
		semverBump = FinalBump{}
	case "pre-release":
		semverBump = PreReleaseBump{
			Pre: preStr,
		}
	}

	if semverBump == nil {
		semverBump = IdentityBump{}
	}

	return semverBump
}
