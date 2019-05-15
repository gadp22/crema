// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Data map[string]interface{}

// ReadDbConfigFile reads database configuration inside
// the config file db.json
func ReadDbConfigFile() map[string]string {
	conf, _ := ioutil.ReadFile("./conf/db.json")

	err := json.Unmarshal(conf, &Data)

	HandleError(err)

	return interfaceToMapStringString(Data["db"])
}

func interfaceToMapStringString(inter interface{}) map[string]string {
	mapStringString := make(map[string]string)
	mapStringInterface := inter.(map[string]interface{})

	for key, value := range mapStringInterface {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapStringString[strKey] = strValue
	}

	return mapStringString
}
