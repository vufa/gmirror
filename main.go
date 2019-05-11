// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/countstarlight/gmirror/cmd"
	"github.com/countstarlight/gmirror/modules/setting"
	"github.com/countstarlight/gmirror/modules/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.MaxConcurrency = runtime.NumCPU() * 2
}

func main() {
	app := cli.NewApp()
	app.Name = "gmirror"
	app.Version = version.Version.String()
	app.Usage = "mirror between git repositories"
	app.Action = cmd.Mirror
	app.Flags = cmd.Flags
	app.Before = cmd.Before
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("[gmirror]%s", err.Error())
	}
}
