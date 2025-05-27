package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func GetType(variable interface{}) string {
	return reflect.TypeOf(variable).String()
}

func Concatenate(vars ...interface{}) string {
	var sb strings.Builder
	for _, v := range vars {
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	return sb.String()
}

func Hash(input string, salt string) string {
	saltedInput := input[:len(input)/2] + salt + input[len(input)/2:]
	hash := sha256.Sum256([]byte(saltedInput))
	return hex.EncodeToString(hash[:])
}
