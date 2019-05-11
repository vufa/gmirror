// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"fmt"
	"github.com/countstarlight/gmirror/modules/setting"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

func Pull(r *git.Repository, repo string, repoRef string, repoAuth transport.AuthMethod) error {
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Pull using default options
	// If authentication required pull using authentication
	var reference plumbing.ReferenceName
	reference = plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", repoRef))

	err = w.Pull(&git.PullOptions{
		RemoteName:    setting.OriginRemoteName,
		ReferenceName: reference,
		Auth:          repoAuth,
	})
	if err != nil {
		switch err {
		case transport.ErrEmptyRemoteRepository:
			return fmt.Errorf("origin repository is empty")
		default:
			return err
		}
	}

	// Print the latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		return err
	}

	_, err = r.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	return nil
}
