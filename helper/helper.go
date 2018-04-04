package helper

import (
	"encoding/json"
	"log"
)

// ToStruct use for converter struct from byte
func ToStruct(source interface{}, dest interface{}) {
	byteSource, errSource := json.Marshal(source)
	CheckError("error marshal source struct", errSource)
	json.Unmarshal(byteSource, dest)
}

// CheckError global
func CheckError(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	}
}
