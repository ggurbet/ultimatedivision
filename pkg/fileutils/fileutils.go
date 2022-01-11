// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package fileutils

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// ReadLine reads certain line from the file.
func ReadLine(file *os.File, lineNum int) (string, error) {
	var line int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++

		if line == lineNum {
			return strings.TrimSpace(scanner.Text()), scanner.Err()
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", io.EOF
}

// CountLines counts lines from file.
func CountLines(r io.Reader) (int, error) {
	var line int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return line, nil
}

// ApplicationDir returns best base directory for specific OS.
func ApplicationDir(subdir ...string) string {
	for i := range subdir {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			subdir[i] = strings.Title(subdir[i])
		} else {
			subdir[i] = strings.ToLower(subdir[i])
		}
	}

	var appdir string

	home := os.Getenv("HOME")
	//
	switch runtime.GOOS {
	case "windows":
		// Windows standards: https://msdn.microsoft.com/en-us/library/windows/apps/hh465094.aspx?f=255&MSPPError=-2147217396
		for _, env := range []string{"AppData", "AppDataLocal", "UserProfile", "Home"} {
			val := os.Getenv(env)
			if val != "" {
				appdir = val
				break
			}
		}
	case "darwin":
		// Mac standards: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html
		appdir = filepath.Join(home, "Library", "Application Support")
	case "linux":
		fallthrough
	default:
		// Linux standards: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
		appdir = os.Getenv("XDG_DATA_HOME")
		if appdir == "" && home != "" {
			appdir = filepath.Join(home, ".local", "share")
		}
	}

	return filepath.Join(append([]string{appdir}, subdir...)...)
}
