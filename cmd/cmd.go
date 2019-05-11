// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/countstarlight/gmirror/modules/setting"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"runtime"
	"strings"
)

var (
	origin   string
	username string
	target   string
	token    string
)

var Flags = []cli.Flag{
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
		Name:        "token, k",
		Usage:       "token or password",
		Destination: &token,
	},
	cli.StringFlag{
		Name:        "username, u",
		Usage:       "username",
		Destination: &username,
	},
}

func Mirror(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		setting.DebugMode = true
		// Set logrus format
		// Print file name and line code
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "15:04:05",
			// Show colorful on windows
			ForceColors:   true,
			FullTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				repopath := fmt.Sprintf("%s/src/github.com/countstarlight/gmirror/", os.Getenv("GOPATH"))
				filename := strings.Replace(f.File, repopath, "", -1)
				r := strings.Split(f.Function, ".")
				return fmt.Sprintf("%s()", r[len(r)-1]), fmt.Sprintf("%s:%d", filename, f.Line)
			},
		})
		logrus.Infof("Running in debug mode")
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "15:04:05",
			// Show colorful on windows
			ForceColors:   true,
			FullTimestamp: true,
		})
	}
	return nil
}

func Before(c *cli.Context) error { return nil }
