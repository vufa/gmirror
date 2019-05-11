// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/countstarlight/gmirror/modules/git"
	"github.com/countstarlight/gmirror/modules/setting"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"runtime"
	"strings"
)

var SyncFlags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "GMIRROR_DEBUG",
		Name:   "debug, d",
		Usage:  "running gmirror in debug mode",
	},
	cli.StringFlag{
		Name: "origin, o",
		//Value: "https://github.com/countstarlight/gmirror.git",
		Usage:       "origin repo",
		Destination: &origin,
	},
	cli.StringFlag{
		Name: "target, t",
		//Value: "https://github.com/countstarlight/gmirror.git",
		Usage:       "target repo",
		Destination: &target,
	},
	cli.StringFlag{
		EnvVar:      "GMIRROR_TOKEN",
		Name:        "token, k",
		Usage:       "token or password",
		Destination: &token,
	},
	cli.StringFlag{
		EnvVar:      "GMIRROR_USERNAME",
		Name:        "username, u",
		Usage:       "username",
		Destination: &username,
	},
}

func Sync(ctx *cli.Context) error {
	if ctx.Bool("debug") || setting.DebugMode {
		setting.DebugMode = true
		// Set logrus format
		// Print file name and line code
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "15:04:05",
			// Display colorful log on windows
			ForceColors:   true,
			FullTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := strings.Replace(f.File, setting.RootPath+"/", "", -1)
				r := strings.Split(f.Function, ".")
				return fmt.Sprintf("%s()", r[len(r)-1]), fmt.Sprintf("%s:%d", filename, f.Line)
			},
		})
		logrus.Infof("Running in debug mode")
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "15:04:05",
			// Display colorful log on windows
			ForceColors:   true,
			FullTimestamp: true,
		})
	}
	logrus.Infof("Gmirror v%s", setting.AppVersion)
	if origin == "" && target == "" {
		logrus.Fatalf("You should set origin and target repo, see 'gmirror help sync' or 'gmirror sync --help'")
	}
	if username == "" && token == "" {
		logrus.Fatal("You should set username or token, see 'gmirror help sync' or 'gmirror sync --help'")
	}
	// Get repo auth
	originAuth, err := git.GetAuth(origin, username, token)
	if err != nil {
		logrus.Fatal(err)
	}
	targetAuth, err := git.GetAuth(target, username, token)
	if err != nil {
		logrus.Fatal(err)
	}

	// Check target repo
	err = git.ValidateRepo(target, targetAuth)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(originAuth)
	return nil
}
