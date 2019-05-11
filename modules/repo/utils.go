// Copyright 2019 Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"strings"
)

// ExtractPath return repo path
func ExtractPath(repo string) (string, error) {
	endpoint, err := transport.NewEndpoint(repo)
	if err != nil {
		return "", err
	}

	path := strings.TrimSuffix(endpoint.Path, ".git")

	// Check for missing path separator
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	host := endpoint.Host
	filePath := host + path

	return filePath, nil
}
