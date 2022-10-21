package utils

import (
	"encoding/json"
	"fmt"
)

func Obj2Str(t any) string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println("err>>", err)
		return ""
	}
	return string(bytes)
}
