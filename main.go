// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/countstarlight/gmirror/cmd"
	"github.com/countstarlight/gmirror/modules/setting"
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

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "show gmirror version",
	}
	app := cli.NewApp()
	app.Name = "gmirror"
	app.Version = setting.AppVersion
	app.Usage = "mirror between git repositories"
	// Global flags
	app.Flags = cmd.Flags
	app.Commands = cmd.Commands
	app.Before = cmd.Before
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("[gmirror]%s", err.Error())
	}
}
