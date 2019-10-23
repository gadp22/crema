// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

// HandleError calls crema.PrintfError(). It then calls panic() which represents
// a go panic to stop the ordinary flow of control and begins panicking.
func HandleError(err error) {
	if err != nil {
		LogPrintfError(err.Error())
	}
}
