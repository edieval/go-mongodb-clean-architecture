package tools

import (
	"encoding/json"
	"log"
)

func JsonToString(obj interface{}) string {
	out, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	log.Print(string(out))
	return string(out)
}
