// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// PopulateSingleValueQueries populates query strings from the incoming requests
// TO DO: documentation will be updated soon
func PopulateSingleValueQueries(r *http.Request, conditions map[string]string) {
	r.ParseForm()
	request := r.Form

	for key, val := range request {
		conditions[key] = val[0]
	}
}

// PopulateParams populates the URI segment params from the incoming requests
// TO DO: documentation will be updated soon
func PopulateParams(params map[string]string, conditions map[string]string) {
	for key, val := range params {
		conditions[key] = val
	}
}

// PopulateRequestBody populates the body from the incoming requests
// TO DO: documentation will be updated soon
func PopulateRequestBody(r *http.Request, conditions map[string]string) {
	body, err := ioutil.ReadAll(r.Body)
	if len(body) > 0 {
		var test interface{}
		err = json.Unmarshal(body, &test)

		HandleError(err)

		for key, val := range interfaceToMapStringString(test.(map[string]interface{})) {
			conditions[key] = val
		}
	}
}
