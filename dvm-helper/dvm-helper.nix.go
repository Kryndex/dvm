// +build !windows

package main

import (
	"os"
	"path/filepath"

	"github.com/howtowhale/dvm/dvm-helper/internal/downloader"
)

const binaryFileExt string = ""

func upgradeSelf(version string) {
	d := downloader.New(getDebugLogger())

	binaryURL := buildDvmReleaseURL(version, dvmOS, dvmArch, "dvm-helper")
	binaryPath := filepath.Join(dvmDir, "dvm-helper", "dvm-helper")
	err := d.DownloadFileWithChecksum(binaryURL, binaryPath)
	if err != nil {
		die("", err, retCodeRuntimeError)
	}

	scriptURL := buildDvmReleaseURL(version, "dvm.sh")
	scriptPath := filepath.Join(dvmDir, "dvm.sh")
	err = d.DownloadFile(scriptURL, scriptPath)
	if err != nil {
		die("", err, retCodeRuntimeError)
	}
}

func getCleanPathRegex() string {
	versionDir := getVersionsDir()
	return versionDir + `/(\d+\.\d+\.\d+|edge):`
}

func validateShellFlag() {
	// we don't care about the shell flag on non-Windows platforms
}

func getUserHomeDir() string {
	return os.Getenv("HOME")
}
