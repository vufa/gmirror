// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import (
	"github.com/countstarlight/gmirror/modules/com"
	"github.com/countstarlight/gmirror/modules/version"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"sync"

	"github.com/go-ini/ini"
)

var (
	//App settings
	AppPath    string
	RootPath   string
	AppVersion string

	//Global setting objects
	Cfg       *ini.File
	DebugMode bool
	IsWindows bool
	ConfFile  string

	// Log settings
	LogRootPath string

	//Maximum number of upload process at same time
	MaxConcurrency int

	//Control multiple goroutines create directory
	MkLock *sync.Mutex
)

// execPath returns the executable path.
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func init() {
	IsWindows = runtime.GOOS == "windows"

	var err error
	if AppPath, err = execPath(); err != nil {
		logrus.Fatalf("Fail to get app path: %s\n", err.Error())
	}

	// Note: we don't use path.Dir here because it does not handle case
	//	which path starts with two "/" in Windows: "//psf/Home/..."
	AppPath = strings.Replace(AppPath, "\\", "/", -1)

	// Get Root path
	RootPath, err = WorkDir()
	if err != nil {
		logrus.Fatalf("Fail to get work directory: %s", err.Error())
	}
	//Control multiple goroutines create directory
	MkLock = new(sync.Mutex)

	AppVersion = version.Version.String()
}

// WorkDir returns absolute path of work directory.
func WorkDir() (string, error) {
	wd := os.Getenv("GMIRROR_WORK_DIR")
	if len(wd) > 0 {
		return wd, nil
	}

	i := strings.LastIndex(AppPath, "/")
	if i == -1 {
		return AppPath, nil
	}
	return AppPath[:i], nil
}

func forcePathSeparator(path string) {
	if strings.Contains(path, "\\") {
		logrus.Fatalf("Do not use '\\' or '\\\\' in paths, instead, please use '/' in all places")
	}
}

// NewContext initializes configuration context.
// NOTE: do not print any log except error.
func NewContext() {
	ConfFile = path.Join(RootPath, "conf/app.ini")

	//Cfg, err = ini.Load("conf/example_app.ini")
	Cfg, err := ini.Load(ConfFile)
	if err != nil {
		logrus.Fatalf("Fail to parse %s: %s", ConfFile, err.Error())
	}

	Cfg.NameMapper = ini.AllCapsUnderscore

	// Load log config
	sec := Cfg.Section("log")
	LogRootPath = sec.Key("ROOT_PATH").MustString(path.Join(RootPath, "log"))
	forcePathSeparator(LogRootPath)
}

func newLogService() {
	if DebugMode {
		logrus.SetOutput(os.Stdout)
	} else {
		logPath := path.Join(LogRootPath, "candy.log")
		if err := os.MkdirAll(path.Dir(logPath), os.ModePerm); err != nil {
			logrus.Fatalf("Fail to create log directory '%s': %s", path.Dir(logPath), err.Error())
		}
		// Create log file if not exist
		if !com.IsFile(logPath) {
			_, err := os.Create(logPath)
			if err != nil {
				logrus.Fatalf("Fail to create log file: %s", err.Error())
			}
		}
		f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			logrus.Fatalf("Fail to open log file: %s", err.Error())
		}
		logrus.SetOutput(f)
	}
}

func NewServices() {
	newLogService()
}
