// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"net/http"
)

// PopulateSingleValueQueries populates query strings from the incoming requests
// TO DO: documentation will be updated soon
func PopulateSingleValueQueries(r *http.Request, conditions map[string]string) map[string]string {
	r.ParseForm()
	request := r.Form

	for key, val := range request {
		conditions[key] = val[0]
	}

	return conditions
}

// PopulateParams populates the URI segment params from the incoming requests
// TO DO: documentation will be updated soon
func PopulateParams(params map[string]string, conditions map[string]string) map[string]string {
	for key, val := range params {
		conditions[key] = val
	}

	return conditions
}
