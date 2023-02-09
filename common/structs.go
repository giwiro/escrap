package common

import "encoding/json"

func StructToByte(obj interface{}) (b []byte, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &b) // Convert to byte
	return
}

func StructToMap(obj interface{}) (m map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &m) // Convert to a map
	return
}
