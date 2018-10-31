package gomodone

import (
	"github.com/muktadiur/gomodtwo/v2"
)

func Version() string {
	return "gomodone v1.0.2"
}

func ModTwoVersion() string {
	return gomodtwo.VersionName()
}
