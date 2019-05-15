// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"net/http"
	"strconv"
)

// GenericHTTPResponse generates a default HTTP response
// TO DO: documentation will be updated soon
func GenericHTTPResponse(code int) Response {
	var response Response

	if code == http.StatusOK {
		response.Status = 1
		response.Message = "Success " + strconv.Itoa(code)
	} else {
		response.Status = 0
		response.Message = "Error " + strconv.Itoa(code)
	}

	return response
}
