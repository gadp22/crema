package crema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Data map[string]interface{}

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
