package gomodone

import (
	"github.com/muktadiur/gomodtwo"
)

func Version() string {
	return "gomodone v1.0.1"
}

func ModTwoVersion() string {
	return gomodtwo.Version()
}
