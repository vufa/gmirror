// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"os"
)

// GetAuth return repo auth
func GetAuth(repo string, username string, token string) (transport.AuthMethod, error) {
	var auth transport.AuthMethod
	endpoint, err := transport.NewEndpoint(repo)
	if err != nil {
		return nil, err
	}
	switch {
	case endpoint.Protocol == "ssh":
		user := os.Getenv("USER")
		auth, err = ssh.NewSSHAgentAuth(user)
		if err != nil {
			return nil, err
		}
	case endpoint.Protocol == "https":
		auth = &http.BasicAuth{
			Username: username,
			Password: token,
		}
	default:
		auth = nil
	}

	return auth, nil
}

// ValidateRepo check repo validate or not
func ValidateRepo(repo string, targetAuth transport.AuthMethod) error {
	// Create a temporary repository in memory
	r, err := git.Init(memory.NewStorage(), nil)
	if err != nil {
		return err
	}

	// Add a new remote, with the default fetch refspec
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: git.DefaultRemoteName,
		URLs: []string{repo},
	})
	if err != nil {
		return err
	}

	// Fetch using the new remote
	err = r.Fetch(&git.FetchOptions{
		RemoteName: git.DefaultRemoteName,
		Auth:       targetAuth,
	})
	if err != nil {
		return err
	}

	return nil
}
