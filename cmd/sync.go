// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/countstarlight/gmirror/modules/repo"
	"github.com/countstarlight/gmirror/modules/setting"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"log"
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
	originAuth, err := repo.GetAuth(origin, username, token)
	if err != nil {
		logrus.Fatalf("Get origin repo auth failed: %s\n", err.Error())
	}
	targetAuth, err := repo.GetAuth(target, username, token)
	if err != nil {
		logrus.Fatalf("Get target repo auth failed: %s\n", err.Error())
	}

	// Check origin repo
	err = repo.ValidateRepo(origin, originAuth)
	if err != nil {
		logrus.Fatalf("Check origin repo failed: %s\n", err.Error())
	}
	path, err := repo.ExtractPath(origin)
	if err != nil {
		logrus.Fatal()
	}
	// Initialize non bare repo
	r, err := git.PlainInit(path, false)
	if err != git.ErrRepositoryAlreadyExists && err != nil {
		logrus.Fatal("Init repo failed: %s\n", err.Error())
	}

	// Open repo, if initialized
	r, err = git.PlainOpen(path)
	if err != nil {
		log.Println(err)
	}

	// Add origin repo remote
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: setting.OriginRemoteName,
		URLs: []string{origin},
	})
	if err != nil {
		if err == git.ErrRemoteNotFound {
			logrus.Fatalf("remote [%s] not found", setting.OriginRemoteName)
		} else {
			logrus.Fatalf("Create remote [%s] failed: %s", setting.OriginRemoteName, err.Error())
		}
	}
	logrus.Infof("Pulling %s...\n", origin)
	err = repo.Pull(r, origin, "master", originAuth)
	if err != nil {
		logrus.Fatalf("Pull repo failed: %s", err.Error())
	}
	logrus.Infof("Pulled: %s...\n", origin)
	logrus.Infof("Pushing to %s...\n", target)
	err = repo.Push(r, target, "master", "master", targetAuth)
	if err != nil {
		switch err {
		case git.NoErrAlreadyUpToDate:
			logrus.Warnf("Target repo already up-to-date")
		default:
			logrus.Fatalf("Push repo failed: %s", err.Error())
		}
	}
	logrus.Info("Repository successfully synced")
	return nil
}
