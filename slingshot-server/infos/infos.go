package infos

import (
	_ "embed"
)

//go:embed version.txt
var version []byte

// GetVersion returns the current version
func GetVersion() string {
	return string(version)
}

//go:embed usage.txt
var usage []byte
var Usage = `Slingshot ` + string(version) + "\n\r\n\r" + string(usage)

//go:embed about.txt
var about []byte
var About = `Slingshot ` + string(version) + "\n\r\n\r" + string(about)

//go:embed help.txt
var help []byte
var Help = `Slingshot ` + string(version) + "\n\r\n\r" + string(help)