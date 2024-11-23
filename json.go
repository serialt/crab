package crab

import "encoding/json"

func Marshal(data any) (str string) {
	bt, _ := json.Marshal(data)
	return string(bt)
}
