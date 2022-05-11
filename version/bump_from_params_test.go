package version_test

import (
	"fmt"

	. "github.com/bitex-la/semver-resource/version"
	"github.com/blang/semver"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BumpForParams", func() {
	var (
		version semver.Version

		bumpParam string
		preParam  string
	)

	BeforeEach(func() {
		version = semver.Version{
			Major: 1,
			Minor: 2,
			Patch: 3,
		}

		bumpParam = ""
		preParam = "pre"
	})

	JustBeforeEach(func() {
		version = BumpFromParams(bumpParam, preParam).Apply(version)
	})

	for bump, result := range map[string]string{
		"":            "1.2.3",
		"final":       "1.2.4",
		"pre-release": "1.2.4-pre.1",
	} {
		bumpLocal := bump
		resultLocal := result

		Context(fmt.Sprintf("when bumping %s", bumpLocal), func() {
			BeforeEach(func() {
				bumpParam = bumpLocal
			})

			It("bumps to "+resultLocal, func() {
				Expect(version.String()).To(Equal(resultLocal))
			})
		})
	}

	Context("when bumping from a prerelease", func() {
		BeforeEach(func() {
			version.Pre = []semver.PRVersion{
				{VersionStr: "pre"},
				{VersionNum: 1, IsNum: true},
			}
		})

		for bump, result := range map[string]string{
			"":            "1.2.3-pre.1",
			"final":       "1.2.3",
			"pre-release": "1.2.3-pre.2",
		} {
			bumpLocal := bump
			resultLocal := result

			Context(fmt.Sprintf("when bumping %s", bumpLocal), func() {
				BeforeEach(func() {
					bumpParam = bumpLocal
				})

				It("bumps to "+resultLocal, func() {
					Expect(version.String()).To(Equal(resultLocal))
				})
			})
		}
	})
})
