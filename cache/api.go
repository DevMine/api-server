// Copyright 2014-2015 The DevMine authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"database/sql"
)

// loadFeaturesNames loads the map of features names.
func loadFeaturesNames(db *sql.DB) error {
	feats := make(map[string]struct{})

	rows, err := db.Query(`SELECT features.name FROM features ORDER BY features.name ASC`)
	if err != nil {
		return err
	}

	for rows.Next() {
		var f string
		if err := rows.Scan(&f); err != nil {
			return err
		}

		feats[f] = struct{}{}
	}

	featuresNames = feats

	return nil
}
