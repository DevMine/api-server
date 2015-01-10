// Copyright 2014-2015 The DevMine authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stats handles /stats... routes.
package stats

import (
	"net/http"

	"github.com/DevMine/api-server/model"
	"github.com/DevMine/api-server/srv/context"
	"github.com/DevMine/api-server/util/json"
)

// Index handles "/stats" route.
func Index(c *context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	var stats model.Stats

	err = c.DB.QueryRow(`SELECT COUNT(users.id) FROM users`).Scan(&stats.UsersCount)
	if err != nil {
		panic(err)
	}

	err = c.DB.QueryRow(`SELECT COUNT(repositories.id) FROM repositories`).Scan(
		&stats.RepositoriesCount)
	if err != nil {
		panic(err)
	}

	err = c.DB.QueryRow(`SELECT COUNT(features.id) FROM features`).Scan(
		&stats.FeaturesCount)
	if err != nil {
		panic(err)
	}

	err = c.DB.QueryRow(`SELECT COUNT(gh_users.id) FROM gh_users`).Scan(
		&stats.GhUsersCount)
	if err != nil {
		panic(err)
	}

	err = c.DB.QueryRow(`SELECT COUNT(gh_organizations.id) FROM gh_organizations`).Scan(
		&stats.GhOrganizationsCount)
	if err != nil {
		panic(err)
	}

	err = c.DB.QueryRow(`SELECT COUNT(gh_repositories.id) FROM gh_repositories`).Scan(
		&stats.GhRepositoriesCount)
	if err != nil {
		panic(err)
	}

	w.Write(json.MarshalIndentPanic(stats))
}
