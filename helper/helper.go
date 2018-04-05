package helper

import (
	"encoding/json"
	"log"
)

// ToStruct use for converter struct from byte
func ToStruct(source interface{}, dest interface{}) {
	byteSource, errSource := json.Marshal(source)
	CheckError("error marshal source struct", errSource)
	errUnmarshal := json.Unmarshal(byteSource, dest)
	CheckError("failed umarshal converter struct", errUnmarshal)
}

// CheckError global
func CheckError(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	}
}
