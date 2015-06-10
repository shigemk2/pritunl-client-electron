package utils

import (
	"github.com/dropbox/godropbox/errors"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"os"
)

func GetRootDir() (pth string) {
	pth, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	return
}

func GetTempDir() (pth string, err error) {
	if runtime.GOOS == "windows" {
		pth = filepath.Join("C:", "ProgramData", "Pritunl")
	} else {
		pth = filepath.Join(string(filepath.Separator), "tmp", "pritunl")
	}

	err = os.MkdirAll(pth, 0700)
	if err != nil {
		err = &IoError{
			errors.Wrap(err, "utils: Failed to create temp directory"),
		}
		return
	}

	return
}

func UpdateAdapters() (adapUsed int, adapTotal int, err error) {
	if runtime.GOOS == "linux" {
		adapUsed = 0
		adapTotal = 100
		return
	}

	output, err := exec.Command("ipconfig", "/all").Output()
	if err != nil {
		err = &CommandError{
			errors.Wrap(err, "utils: Update tuntap adapters failed"),
		}
		return
	}

	adap := false
	adapDisc := false
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		line = strings.Trim(line, "\r")

		if line == "" {
			if adap {
				adapTotal += 1
				if !adapDisc {
					adapUsed += 1
				}
			}
			adap = false
			adapDisc = false
		} else if strings.Contains(line, "TAP-Windows Adapter V9") {
			adap = true
		} else if strings.Contains(line, "Media disconnected") {
			adapDisc = true
		}
	}

	return
}