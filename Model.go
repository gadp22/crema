// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

// Response represents the base of a HTTP response message
// TO DO: documentation will be updated soon
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// GenericResponse represents the default response message
// TO DO: documentation will be updated soon
type GenericResponse struct {
	Response
	Data interface{} `json:"data"`
}
