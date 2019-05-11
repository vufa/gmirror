// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"fmt"
	"github.com/countstarlight/gmirror/modules/setting"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

func Push(r *git.Repository, target string, originRef string, targetRef string, targetAuth transport.AuthMethod) error {
	// Validate target URL
	err := ValidateRepo(target, targetAuth)
	if err != nil {
		return err
	}

	// Add a target remote
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: setting.TargetRemoteName,
		URLs: []string{target},
	})
	if err != nil {
		switch err {
		case git.ErrRemoteNotFound:
			return fmt.Errorf("remote not found")
		default:
			return err
		}
	}

	// Push using default options
	// If authentication required push using authentication
	referenceList := append([]config.RefSpec{}, config.RefSpec(originRef+":"+targetRef))
	err = r.Push(&git.PushOptions{
		RemoteName: setting.TargetRemoteName,
		RefSpecs:   referenceList,
		Auth:       targetAuth,
	})
	if err != nil {
		return err
	}

	return nil
}
