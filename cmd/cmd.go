// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/countstarlight/gmirror/modules/setting"
	"github.com/urfave/cli"
)

var (
	origin   string
	username string
	target   string
	token    string
)

var Commands = []cli.Command{
	{
		Name:    "sync",
		Aliases: []string{"s"},
		Usage:   "sync git repositories",
		Flags:   SyncFlags,
		Action:  Sync,
	},
}

var Flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar:      "GMIRROR_DEBUG",
		Name:        "debug, d",
		Usage:       "running gmirror in debug mode",
		Destination: &setting.DebugMode,
	},
}

func Before(c *cli.Context) error { return nil }
